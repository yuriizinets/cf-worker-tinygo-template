package main

import (
	"cloudflare.worker/cloudflare"
)

func entrypoint(r cloudflare.Request) cloudflare.Response {
	return cloudflare.Response{
		Body: []byte(`{"foo":"bar"}`),
	}
}

func main() {
	cloudflare.ListenAndServe(entrypoint)
}
