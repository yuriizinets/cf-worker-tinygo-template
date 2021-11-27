

dev:
	GOOS=js GOARCH=wasm tinygo build -target wasm -o cloudflare/worker/module.wasm .
	cd cloudflare && \
	wrangler dev

publish:
	GOOS=js GOARCH=wasm tinygo build -target wasm -o cloudflare/worker/module.wasm .
	cd cloudflare && \
	wrangler publish
