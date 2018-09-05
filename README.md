# ginkgomod
Ginkgo on go modules

## Environment

```console
$ uname -a
Linux host-vm 4.9.0-8-amd64 #1 SMP Debian 4.9.110-3+deb9u3 (2018-08-19) x86_64 GNU/Linux

$ go version
go version go1.11 linux/amd64
```

## Run

```console
# Enable go modules if inside GOPATH
$ export GO111MODULE=on

$ go mod download
$ go get github.com/onsi/ginkgo/ginkgo
```

```console
$ ginkgo .
Failed to compile ginkgomod:

go test golang_org/x/crypto/cryptobyte/asn1: open /usr/local/go/pkg/linux_amd64/vendor/golang_org/x/crypto/cryptobyte/asn1.a: permission denied
go test golang_org/x/net/dns/dnsmessage: open /usr/local/go/pkg/linux_amd64/vendor/golang_org/x/net/dns/dnsmessage.a: permission denied
go test golang_org/x/crypto/curve25519: open /usr/local/go/pkg/linux_amd64/vendor/golang_org/x/crypto/curve25519.a: permission denied
go test golang_org/x/text/transform: open /usr/local/go/pkg/linux_amd64/vendor/golang_org/x/text/transform.a: permission denied
go test golang_org/x/crypto/poly1305: open /usr/local/go/pkg/linux_amd64/vendor/golang_org/x/crypto/poly1305.a: permission denied
go test golang_org/x/crypto/internal/chacha20: open /usr/local/go/pkg/linux_amd64/vendor/golang_org/x/crypto/internal/chacha20.a: permission denied
go test golang_org/x/net/http2/hpack: open /usr/local/go/pkg/linux_amd64/vendor/golang_org/x/net/http2/hpack.a: permission denied
go test golang_org/x/text/unicode/bidi: open /usr/local/go/pkg/linux_amd64/vendor/golang_org/x/text/unicode/bidi.a: permission denied

Ginkgo ran 1 suite in 953.207656ms
```
