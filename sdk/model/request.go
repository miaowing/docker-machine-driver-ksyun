package model

import "net/url"

type Request struct {
	Method  string
	Uri     string
	Querys  url.Values
	Headers map[string]string
	Body    []byte
}
