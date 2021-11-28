

dev:
	mkdir -p cloudflare/worker
	GOOS=js GOARCH=wasm tinygo build -target wasm -o cloudflare/worker/module.wasm .
	cd cloudflare && \
	wrangler dev

publish:
	mkdir -p cloudflare/worker
	GOOS=js GOARCH=wasm tinygo build -target wasm -o cloudflare/worker/module.wasm .
	cd cloudflare && \
	wrangler publish
