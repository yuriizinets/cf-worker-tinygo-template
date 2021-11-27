package cloudflare

import (
	"net/url"
	"syscall/js"
)

type Request struct {
	URL     *url.URL
	Headers map[string]string
	Body    []byte
}

type Response struct {
	Body []byte
}

type Entrypoint func(Request) Response

var entrypoint Entrypoint

func bridge(_ js.Value, input []js.Value) interface{} {
	// Unpack values
	_url, _ := url.Parse(input[0].Get("url").String())
	// Request processing
	resp := entrypoint(Request{
		URL:     _url,
		Headers: map[string]string{},
		Body:    []byte{},
	})
	// Pack values to serializable types
	body := js.Global().Get("Uint8Array").New(len(resp.Body))
	js.CopyBytesToJS(body, resp.Body)
	// Return
	input[1].Invoke(map[string]interface{}{
		"Body": body,
	})
	return nil
}

func ListenAndServe(e Entrypoint) {
	entrypoint = e
	js.Global().Set("bridge", js.FuncOf(bridge))
	<-make(chan interface{})
}
