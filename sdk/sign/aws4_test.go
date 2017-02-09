package sign

import (
	"fmt"
	"net/url"
	"testing"
)

func TestAws4Sign(t *testing.T) {
	keys := &Keys{AccessKey: "AccessKey", SecretKey: "SecretKey"}
	service := &Service{Name: "iam", Region: "cn-beijing-6"}

	querys := (&url.URL{}).Query()
	headers := map[string]string{"host": "iam.api.ksyun.com"}

	querys.Set("Action", "ListUsers")
	querys.Set("Version", "2015-11-01")
	querys.Set("X-Amz-Algorithm", GetAlgorithm())
	querys.Set("X-Amz-Credential", GetCredential(keys, service))
	querys.Set("X-Amz-Date", GetDate())
	querys.Set("X-Amz-SignedHeaders", GetSignHeaders(headers))

	signature := GetSignature(keys, service, &Request{
		Method:  "GET",
		Uri:     "/",
		Headers: headers,
		Querys:  querys})

	fmt.Println("the querystring is ", querys.Encode())
	fmt.Println("the headers is ", headers)
	fmt.Println("the signature is ", signature)
}
