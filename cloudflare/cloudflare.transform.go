package cloudflare

import (
	"net/url"
	"syscall/js"
)

func decodeHeaders() {
	// TODO: Implement headers
}

func decodeRequest(v js.Value) *Request {
	// Unpack values
	_url, _ := url.Parse(v.Get("url").String())
	// Compose and return
	return &Request{
		URL:     _url,
		Headers: map[string]string{},
		Body:    []byte{},
	}
}

func encodeResponse(r *Response) map[string]interface{} {
	// Pack values to serializable types
	body := js.Global().Get("Uint8Array").New(len(r.Body))
	js.CopyBytesToJS(body, r.Body)
	return map[string]interface{}{
		"Body": body,
	}
}
