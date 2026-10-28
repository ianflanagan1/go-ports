# go-ports

The plain values that ports exchange, shared by every project on [go-baseline](https://github.com/ianflanagan1/go-baseline). It's a `pure-library`: nothing here does I/O, so a project's core can build requests and queries and read responses and rows, and its shell can declare ports in these types. [go-adapters](https://github.com/ianflanagan1/go-adapters) implements those ports.

| Package | Values | Adapter in go-adapters |
|---|---|---|
| `httpmsg` | `Request`, `Response` | `httpsend` |
| `sqlmsg` | `Query`, `Rows`, `Result`, `Value`, and the `Tx` port for transactions | `sqlquery` |

## Usage

```go
import "github.com/ianflanagan1/go-ports/httpmsg"
```

## Architecture

This is a `pure-library` module. `go tool go-baseline map` lists each package's role. The rules are in `BASELINE.md`.
