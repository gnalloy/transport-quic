# Usage

[简体中文](usage.zh-CN.md) | [Docs Index](README.md)

## Requirements

- Go 1.25 or newer, matching the module `go` directive.
- A Gnalloy application, recipe, example, or benchmark harness that owns lifecycle and deployment configuration.
- Standalone module verification should set `GOWORK=off` so the module is tested through its published dependency graph.

## Install
```bash
go get gnalloy.org/transport-quic@dev
```

## Import
```go
import "gnalloy.org/transport-quic"
```

## Integration Pattern
- Address, listener, dialer, buffer allocator, event loop, and channel initializer choices are part of the transport boundary.
- Platform-specific transports must return explicit unsupported errors rather than silently falling back.
- Privileged transports such as raw sockets and L2 capture need operating-system capabilities outside the Go module.
- Protocol, TLS, proxy, and observability handlers should be installed through the Channel pipeline.
- QUIC and HTTP/3 paths require TLS 1.3 and a valid ALPN such as `h3`, `doq`, or the WebTransport profile in use.

## API Selection

Use the API inventory to choose the exact constructor or option type for your protocol path:

```bash
go doc gnalloy.org/transport-quic
```

Common current entry points:
- `const DefaultALPN = "gnalloy-quic" ...`
- `const DefaultClientTokenStoreMaxOrigins = 64 ...`
- `var ErrMissingAddress = errors.New("gnalloy/transport/quic: missing address") ...`
- `type Config struct{ ... }`
- `type DialerFunc func(ctx context.Context, addr string, cfg Config) (Connection, error)`
- `type QLogConfig struct{ ... }`
- `const DefaultMaxFrameSize = 1 << 16`
- `var ErrFrameTooLarge = errors.New("gnalloy/transport/quic/application: frame too large") ...`

## Cross-Module Assembly

When multiple Gnalloy repositories are developed together, create a local `go.work` file in your chosen workspace. Keep application-local `replace` directives out of published library modules unless the change is intentionally temporary and never committed.

## Error Handling

Network input, peer behavior, platform capability, and timeout failures must be handled as normal errors. Do not recover protocol correctness by panicking. Return or propagate the module error and close the affected Channel when ownership requires it.
