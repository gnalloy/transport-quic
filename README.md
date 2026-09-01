# transport-quic

[简体中文](README.zh-CN.md) | [Documentation](docs/README.md)

RFC 9000 QUIC transport for Gnalloy backed by quic-go, plus stream and datagram application exchangers.

This module owns an I/O boundary. It creates or adapts Gnalloy Channels for a concrete transport while protocol parsing, business handlers, TLS policy, and observability remain in other modules.

## Status

- Import path: `gnalloy.org/transport-quic`
- Repository: `github.com/gnalloy/transport-quic`
- Default branch: `dev`
- Preview install: `go get gnalloy.org/transport-quic@dev`
- License: Apache-2.0

## Install
```bash
go get gnalloy.org/transport-quic@dev
go doc gnalloy.org/transport-quic
GOWORK=off GOTOOLCHAIN=local go test ./... -count=1
```

## Documentation
- [Overview](docs/overview.md) ([中文](docs/overview.zh-CN.md))
- [Usage](docs/usage.md) ([中文](docs/usage.zh-CN.md))
- [Examples](docs/examples.md) ([中文](docs/examples.zh-CN.md))
- [Configuration](docs/configuration.md) ([中文](docs/configuration.zh-CN.md))
- [Testing and Performance](docs/testing.md) ([中文](docs/testing.zh-CN.md))
- [API Reference](docs/api.md) ([中文](docs/api.zh-CN.md))
- [Notes and Caveats](docs/notes.md) ([中文](docs/notes.zh-CN.md))
- [ADR-001 Module Boundary](docs/decisions/0001-module-boundary.md) ([中文](docs/decisions/0001-module-boundary.zh-CN.md))

## Module Boundary

This repository owns: RFC 9000 QUIC transport for Gnalloy backed by quic-go, plus stream and datagram application exchangers.

It does not absorb neighboring module responsibilities. Core primitives stay in `gnalloy.org/gnalloy`; protocol codecs, transports, handlers, resolvers, examples, and benchmarks stay in their own repositories.

## Packages
- `gnalloy.org/transport-quic` (`quic`)
- `gnalloy.org/transport-quic/application` (`application`)

## Gnalloy Dependencies
- `gnalloy.org/gnalloy`

## Common Integration Pattern
- Address, listener, dialer, buffer allocator, event loop, and channel initializer choices are part of the transport boundary.
- Platform-specific transports must return explicit unsupported errors rather than silently falling back.
- Privileged transports such as raw sockets and L2 capture need operating-system capabilities outside the Go module.
- Protocol, TLS, proxy, and observability handlers should be installed through the Channel pipeline.
- QUIC and HTTP/3 paths require TLS 1.3 and a valid ALPN such as `h3`, `doq`, or the WebTransport profile in use.

## Current Public Entry Points

The generated API reference lists the full public surface. Common constructors or option types currently include:
- `const DefaultALPN = "gnalloy-quic" ...`
- `const DefaultClientTokenStoreMaxOrigins = 64 ...`
- `var ErrMissingAddress = errors.New("gnalloy/transport/quic: missing address") ...`
- `type Config struct{ ... }`
- `type DialerFunc func(ctx context.Context, addr string, cfg Config) (Connection, error)`
- `type QLogConfig struct{ ... }`
- `const DefaultMaxFrameSize = 1 << 16`
- `var ErrFrameTooLarge = errors.New("gnalloy/transport/quic/application: frame too large") ...`

## Verification

```bash
GOWORK=off GOTOOLCHAIN=local go test ./... -count=1
GOWORK=off GOTOOLCHAIN=local go vet ./...
GOWORK=off GOTOOLCHAIN=local go test ./... -run '^$' -bench . -benchmem -count=1
```

For pressure tests, assemble this module with the relevant transport, codec, and handler stack and run the scenario from `gnalloy.org/benchmarks` or `gnalloy.org/examples`. Keep host, operating system, payload, concurrency, warmup, and repetitions in the report.

## Caveats
- This repository is intentionally narrow. Cross-module behavior should be assembled in applications, recipes, examples, or benchmark harnesses.
- Public APIs should remain Go-native and explicit; avoid runtime scanning, hidden global registries, and reflection-heavy behavior in hot paths.
- Treat network input as untrusted. Configure parser limits and return typed errors instead of panics.
- Keep benchmark claims tied to a concrete host, operating system, protocol, payload, concurrency, warmup, and repetition count.
- Transport modules do not parse application protocols by themselves; install the right codec and handler modules in the pipeline.
