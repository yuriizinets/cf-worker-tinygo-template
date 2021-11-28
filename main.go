package main

import (
	"fmt"

	"cloudflare.worker/cloudflare"
	"github.com/google/uuid"
)

func HandlerIndex(r *cloudflare.Request) *cloudflare.Response {
	uuid := uuid.New().String()
	return &cloudflare.Response{
		Body: []byte(fmt.Sprintf(`{"uuid":"%s"}`, uuid)),
	}
}

func main() {
	// Routes
	cloudflare.HandleFunc("/", HandlerIndex)

	// Start
	cloudflare.ListenAndServe()
}
