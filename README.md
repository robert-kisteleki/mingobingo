
# mingobingo - a single Go binary that serves some content

A minimalistic demonstration of:
  * serving static content, embedded into your compiled Go binary
  * serving a JSON based API call (and querying it from your page via fetch())
  * serving websocket-based contents (and querying it from your page via WebSocket())

Not that *minimalistic* means:
* the code is not very elegant, it just does the bare minimum
* there's not much error handling, it will all be just fine!
* the HTML refers to one CSS to show how one can refer to other assets
* the Javascript uses no external libraries

Once you compile / run the program, you can navigate your browser to http://localhost:8080/ to watch all three functions working.

# Licence / contact

(C) Robert Kisteleki

Licensed under MIT license