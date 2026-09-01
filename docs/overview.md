# Overview

[简体中文](overview.zh-CN.md) | [Docs Index](README.md)

## Purpose

RFC 9000 QUIC transport for Gnalloy backed by quic-go, plus stream and datagram application exchangers.

This module owns an I/O boundary. It creates or adapts Gnalloy Channels for a concrete transport while protocol parsing, business handlers, TLS policy, and observability remain in other modules.

## Repository Identity

- Module path: `gnalloy.org/transport-quic`
- GitHub repository: `github.com/gnalloy/transport-quic`
- Default branch: `dev`
- License: Apache-2.0

## Package Map
- `gnalloy.org/transport-quic` (`quic`)
- `gnalloy.org/transport-quic/application` (`application`)

## Direct Gnalloy Dependencies

- `gnalloy.org/gnalloy`

## Direct Dependents in the Current Repository Set

- `gnalloy.org/benchmarks`
- `gnalloy.org/examples`
- `gnalloy.org/resolver-dns-quic`
- `gnalloy.org/transport-http3`
- `gnalloy.org/transport-webtransport`

## Architecture Position

Gnalloy keeps the core small and dependency-light. This repository is a replaceable module around one responsibility, connected through explicit Go packages instead of runtime discovery.

## Compatibility

The public import path is `gnalloy.org/transport-quic`. Until the first stable tag is published, use `@dev` or an explicit pseudo-version selected by your dependency policy.
