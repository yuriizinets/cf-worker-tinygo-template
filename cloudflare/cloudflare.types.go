package cloudflare

import "net/url"

type Request struct {
	URL     *url.URL
	Headers map[string]string
	Body    []byte
}

type Response struct {
	Headers map[string]string
	Body    []byte
}

type Handler func(*Request) *Response
