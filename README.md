# HTTP/1.1 Server in Go

A from-scratch implementation of an HTTP/1.1 server written in Go, built as a learning project to understand how HTTP works under the hood.

## Why this exists

Go's standard library makes it trivial to spin up an HTTP server with `net/http`. This project intentionally avoids that, instead building the protocol handling directly on top of raw TCP connections (`net.Listen`, `net.Conn`) to learn more about Go, and HTTP as a whole.

## Status

**Work in progress.** This is an educational project.

## Prerequisites

- [Go](https://go.dev/dl/) 1.21 or later (adjust to match your `go.mod`)

## Getting Started

Clone the repository and run the server directly with `go run`:

```bash
git clone https://github.com/alex-wanger/http-server.git
cd http-server
go run ./cmd/tcplistener
```

By default the server will start listening for incoming TCP connections and parse them as HTTP/1.1 traffic. Once running, you can test it with `curl`:

```bash
curl -v http://localhost:42069/
```

> Update the port above to match whatever the server is configured to listen on.


## Roadmap

- [ ] Request line and header parsing
- [ ] Request body / chunked transfer-encoding support
- [ ] Response writer with correct status lines and headers
- [ ] Keep-alive / connection management
- [ ] Routing
- [ ] Basic middleware support
- [ ] Test coverage against the HTTP/1.1 spec (RFC 9112)

## References

- [RFC 9112 — HTTP/1.1](https://www.rfc-editor.org/rfc/rfc9112)
- [RFC 9110 — HTTP Semantics](https://www.rfc-editor.org/rfc/rfc9110)

## License

[MIT](LICENSE) — adjust as appropriate for your project.
