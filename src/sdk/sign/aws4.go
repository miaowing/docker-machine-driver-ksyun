package sign

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	iSO8601BasicFormat      = "20060102T150405Z"
	iSO8601BasicFormatShort = "20060102"
	aws4_hmac_sha256        = "AWS4-HMAC-SHA256"
)

// Keys holds a set of Amazon Security Credentials.
type Keys struct {
	AccessKey string
	SecretKey string
}

// Service represents an AWS-compatible service.
type Service struct {
	// Name is the name of the service being used (i.e. iam, etc)
	Name string

	// Region is the region you want to communicate with the service through. (i.e. us-east-1)
	Region string
}

type Request struct {
	Method  string
	Uri     string
	Querys  url.Values
	Headers map[string]string
}

func GetDate() string {
	t := time.Now().UTC()
	return t.Format(iSO8601BasicFormat)
}

func GetShortDate() string {
	t := time.Now().UTC()
	return t.Format(iSO8601BasicFormatShort)
}

func GetCredential(keys *Keys, service *Service) string {
	return keys.AccessKey + "/" + service.creds()
}

func GetAlgorithm() string {
	return aws4_hmac_sha256
}

func GetSignHeaders(headers map[string]string) string {
	i, keys := 0, make([]string, len(headers))
	for k := range headers {
		keys[i] = strings.ToLower(k)
		i++
	}

	sort.Strings(keys)
	w := bytes.NewBufferString("")
	for i, s := range keys {
		if i > 0 {
			w.Write([]byte{';'})
		}
		w.Write([]byte(s))
	}

	return w.String()
}

func GetSignature(keys *Keys, service *Service, request *Request) string {
	keySigned := keys.sign(service)

	return to16(ghmac(keySigned, service.getStringToSign(request).Bytes()))
}

func (k *Keys) sign(s *Service) []byte {
	h := ghmac([]byte("AWS4"+k.SecretKey), []byte(GetShortDate()))
	h = ghmac(h, []byte(s.Region))
	h = ghmac(h, []byte(s.Name))
	h = ghmac(h, []byte("aws4_request"))
	return h
}

func (s *Service) getStringToSign(request *Request) *bytes.Buffer {
	buffer := bytes.NewBufferString("")
	buffer.Write([]byte(GetAlgorithm() + "\n"))
	buffer.Write([]byte(GetDate() + "\n"))
	buffer.Write([]byte(GetShortDate() + "/" + s.Region + "/" + s.Name + "/aws4_request\n"))
	buffer.Write([]byte(to16(gSha256(s.getSignRequestStrBuffer(request).Bytes()))))

	return buffer
}

func (s *Service) getSignRequestStrBuffer(request *Request) *bytes.Buffer {
	buffer := bytes.NewBufferString("")
	buffer.Write([]byte(strings.ToUpper(request.Method) + "\n"))
	buffer.Write([]byte(request.Uri + "\n"))
	buffer.Write([]byte(s.getQueryString(request.Querys).String() + "\n"))
	buffer.Write([]byte(s.getHeaders(request.Headers).String() + "\n\n"))
	buffer.Write([]byte(GetSignHeaders(request.Headers) + "\n"))
	buffer.Write([]byte(to16(gSha256(nil))))

	return buffer
}

func (s *Service) getQueryString(qs url.Values) *bytes.Buffer {
	return bytes.NewBufferString(qs.Encode())
}

func (s *Service) getHeaders(headers map[string]string) *bytes.Buffer {
	i, keys := 0, make([]string, len(headers))
	for key := range headers {
		keys[i] = key
		i++
	}

	sort.Strings(keys)
	w := bytes.NewBufferString("")

	for i := 0; i < len(keys); i = i + 1 {
		w.Write([]byte(strings.ToLower(keys[i]) + ":" + strings.Trim(headers[keys[i]], " ")))
		if i < len(keys)-1 {
			w.Write([]byte("\n"))
		}
	}

	return w
}

func (s *Service) creds() string {
	return GetShortDate() + "/" + s.Region + "/" + s.Name + "/aws4_request"
}

func ghmac(key, data []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

func gSha256(data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func to16(bytes []byte) string {
	return hex.EncodeToString(bytes)
}
