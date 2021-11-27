require('./wasm_exec.js')

addEventListener('fetch', event => {
  event.respondWith(handleRequest(event.request))
})

/**
 * Respond with hello worker text
 * @param {Request} request
 */
function entrypointWrapper(request) {
  const url = new URL(request.url)
  return new Promise((resolve, reject) => {
    try {
      bridge(request, (answer) => {
        resolve(new Response(
          answer.Body,
          {
            status: 200,
          }
        ))
      })
    } catch (e) {
      reject(e)
    }
  })
}

/**
 * Respond with hello worker text
 * @param {Request} request
 */
async function handleRequest(request) {
  const go = new Go()
  const instance = await WebAssembly.instantiate(WASM_MODULE, go.importObject)
  go.run(instance)
  return await entrypointWrapper(request)
  // return new Response(answer, { status: 200, headers: { 'Content-Type': 'text/html' } })
}
