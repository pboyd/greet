# Hello World Web Service — Design

**Date:** 2026-04-23
**Purpose:** Learning exercise to observe downstream effects of renaming a Go module path.

## Goal

Build a minimal Go HTTP service with enough internal structure to make module rename propagation visible and non-trivial.

## Module

`github.com/pboyd/hello`

## Structure

```
go.mod
main.go                  # HTTP server, wires handler + greeting
internal/greeting/       # builds the greeting string
pkg/helloclient/         # thin HTTP client for the service
```

## Components

### `internal/greeting`

Exports a single function:

```go
func Message() string
```

Returns the string `"Hello, World!"`. Private to this module — demonstrates that `internal/` paths break if the module path changes.

### `pkg/helloclient`

Exports a `Client` type:

```go
type Client struct { ... }
func New(baseURL string) *Client
func (c *Client) Hello() (string, error)
```

Makes a GET request to `/` and returns the response body as a string. Publicly importable — external consumers reference it by the full module path.

### `main.go`

Starts an HTTP server on a configurable port (default `:8080`). Registers a single handler at `GET /` that calls `greeting.Message()` and writes it as plain text with status 200.

## Rename Surface

When the module path changes (e.g. `github.com/pboyd/hello` → `github.com/pboyd/greeter`), these locations must be updated:

1. `go.mod` — the `module` directive
2. `main.go` — import of `github.com/pboyd/hello/internal/greeting`
3. Any external consumer of `pkg/helloclient` — their import path changes

## Constraints

- Standard library only — no external dependencies
- No configuration files beyond `go.mod`
