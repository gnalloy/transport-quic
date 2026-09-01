# Examples

[简体中文](examples.zh-CN.md) | [Docs Index](README.md)

## Example 1: Add the Module to an Application

```bash
mkdir gnalloy-app && cd gnalloy-app
go mod init example.com/gnalloy-app
go get gnalloy.org/transport-quic@dev
go doc gnalloy.org/transport-quic
```

## Example 2: Inspect Current Packages

The current source tree exposes these package import paths:
- `gnalloy.org/transport-quic`
- `gnalloy.org/transport-quic/application`

Use `go doc` against the package that matches the behavior you need:

```bash
go doc gnalloy.org/transport-quic
go doc gnalloy.org/transport-quic/application
```

Selected current exported entry points:
- `const DefaultALPN = "gnalloy-quic" ...`
- `const DefaultClientTokenStoreMaxOrigins = 64 ...`
- `var ErrMissingAddress = errors.New("gnalloy/transport/quic: missing address") ...`
- `type ApplicationErrorCode uint64`
- `type CapabilitySet struct{ ... }`
- `type ClientToken = nativequic.ClientToken`
- `const DefaultMaxFrameSize = 1 << 16`
- `var ErrFrameTooLarge = errors.New("gnalloy/transport/quic/application: frame too large") ...`
- `type DatagramExchanger struct{ ... }`
- `type DatagramMatcher func(request []byte, response []byte) bool`
- `type LengthPrefixedCodec struct{ ... }`
- `type Stream = quic.Stream`

## Example 3: Use Executable Tests as Behavioral Examples

Repository tests are executable examples of supported behavior. Start with the selected names below, then read the matching `_test.go` files for complete setup and assertions. See [Testing and Performance](testing.md) for the complete discovered list.

```bash
GOWORK=off GOTOOLCHAIN=local go test ./... -run Test -count=1
```

Selected current test, benchmark, fuzz, and example entry points:
- `TestDatagramExchangerFiltersUnmatchedPayloads`
- `TestDetectNativeSupportReportsRFC9000Boundary`
- `TestDialAddrEarlyValidatesSessionBoundary`
- `TestEvaluateCapabilitiesReportsWebTransportWhenEnabled`
- `TestExternalInteropHandshake`
- `TestLengthPrefixedCodecRejectsOversizedFrame`
- `TestLengthPrefixedCodecRoundTrip`
- `TestListenAddrEarlyValidatesServerBoundary`
- `TestListenDialAddrEarlyUses0RTTAfterSessionResumption`
- `TestListenDialAddrEchoOverRFC9000QUIC`
- `TestListenDialAddrSupportsUnidirectionalStream`
- `TestNewClientTokenStoreRejectsInvalidCapacity`
- `TestNormalizeConfigClonesTLSAndDefaultsRFC9000`
- `TestNormalizeConfigEnables0RTTTokenStore`
- `TestNormalizeConfigEnablesWebTransportPrerequisites`
- `TestNormalizeConfigInstallsQLogTracer`
- `TestNormalizeConfigRejectsInvalidBoundaries`
- `TestQLogFactoryFuncNilSkipsTrace`

## Example 4: Cross-Module Assembly

Direct Gnalloy dependencies for this module:
- `gnalloy.org/gnalloy`

Assembly guidance:
- Use this transport to own the concrete I/O endpoint and connect it to Gnalloy Channel and EventLoop contracts.
- Protocol parsing stays in codec modules and policy stays in handler modules.
- Platform-specific capability checks should happen during application startup and integration tests.

## Example 5: Pressure-Test Harness

For sustained load, wire this module into a scenario under `gnalloy.org/benchmarks` or a runnable client under `gnalloy.org/examples` when the module participates in network traffic. Record host, OS, CPU, Go version, protocol, payload, concurrency, warmup, repetitions, throughput, and p99 latency in the report.
