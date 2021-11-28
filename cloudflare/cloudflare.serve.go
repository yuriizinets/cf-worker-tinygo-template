package cloudflare

import (
	"syscall/js"
)

var handlers = map[string]Handler{}

func HandleFunc(path string, handler Handler) {
	handlers[path] = handler
}

func ListenAndServe() {
	js.Global().Set("bridge", js.FuncOf(bridge))
	<-make(chan interface{})
}
