package cloudflare

import (
	"syscall/js"
)

func bridge(_ js.Value, input []js.Value) interface{} {
	// Decode request
	request := decodeRequest(input[0])
	// Get handler
	handler, ok := handlers["/"+request.URL.RawPath]
	if !ok {
		panic("Handler not found for " + "/" + request.URL.RawPath)
	}
	// Call handler
	resp := handler(request)
	// Pack values to serializable types
	body := js.Global().Get("Uint8Array").New(len(resp.Body))
	js.CopyBytesToJS(body, resp.Body)
	// Return
	input[1].Invoke(encodeResponse(resp))
	return nil
}
