package main

import (
	"cloudflare.worker/cloudflare"
)

func HandlerIndex(r *cloudflare.Request) cloudflare.Response {
	return cloudflare.Response{
		Body: []byte(`{"foo":"bar"}`),
	}
}

func main() {
	// Routes
	cloudflare.HandleFunc("/", HandlerIndex)

	// Start
	cloudflare.ListenAndServe()
}
