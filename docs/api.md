# API Reference

[简体中文](api.zh-CN.md) | [Docs Index](README.md)

This inventory is generated from `go doc -short` for the packages in this repository. It is a quick public-surface map; source files and tests remain the authority for exact semantics.

## Packages

### `gnalloy.org/transport-quic`

Package name: `quic`

```text
const DefaultALPN = "gnalloy-quic" ...
const DefaultClientTokenStoreMaxOrigins = 64 ...
var ErrMissingAddress = errors.New("gnalloy/transport/quic: missing address") ...
type ApplicationErrorCode uint64
type CapabilitySet struct{ ... }
    func EvaluateCapabilities(role EndpointRole, cfg Config) (CapabilitySet, error)
type ClientToken = nativequic.ClientToken
type ClientTokenStore = nativequic.TokenStore
    func NewClientTokenStore(maxOrigins int, tokensPerOrigin int) (ClientTokenStore, error)
type Config struct{ ... }
    func DefaultConfig() Config
    func NormalizeConfig(cfg Config) (Config, error)
type Connection interface{ ... }
    func DialAddr(ctx context.Context, addr string, cfg Config) (Connection, error)
    func DialAddrEarly(ctx context.Context, addr string, cfg Config) (Connection, error)
type ConnectionStats struct{ ... }
type DefaultDialer struct{}
type Dialer interface{ ... }
type DialerFunc func(ctx context.Context, addr string, cfg Config) (Connection, error)
type EarlyDialer interface{ ... }
type EarlyListener interface{ ... }
    func ListenAddrEarly(addr string, cfg Config) (EarlyListener, error)
type EndpointRole uint8
    const EndpointRoleClient EndpointRole = iota + 1 ...
type FeatureCapability struct{ ... }
type FeatureSupport struct{ ... }
type Listener interface{ ... }
    func ListenAddr(addr string, cfg Config) (Listener, error)
type NativeEngine string
    const NativeEngineQUICGo NativeEngine = "quic-go"
type NativeSupport struct{ ... }
    func DetectNativeSupport() NativeSupport
type QLogConfig struct{ ... }
type QLogTraceInfo struct{ ... }
type QLogWriterFactory interface{ ... }
type QLogWriterFactoryFunc func(ctx context.Context, info QLogTraceInfo) (io.WriteCloser, error)
type ReceiveStream interface{ ... }
type SendStream interface{ ... }
type Server struct{ ... }
type State struct{ ... }
type Stream interface{ ... }
type StreamErrorCode uint64
type StreamID int64
type Transport struct{ ... }
    func NewTransport(cfg Config) *Transport
type Version uint32
    const Version1 Version = 0x00000001
```

### `gnalloy.org/transport-quic/application`

Package name: `application`

```text
const DefaultMaxFrameSize = 1 << 16
var ErrFrameTooLarge = errors.New("gnalloy/transport/quic/application: frame too large") ...
type DatagramExchanger struct{ ... }
type DatagramMatcher func(request []byte, response []byte) bool
type LengthPrefixedCodec struct{ ... }
type Stream = quic.Stream
type StreamCodec interface{ ... }
type StreamExchanger struct{ ... }
```
