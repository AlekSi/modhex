# modhex

[![Go Reference](https://pkg.go.dev/badge/github.com/AlekSi/modhex.svg)](https://pkg.go.dev/github.com/AlekSi/modhex)

Go package `modhex` implements hexadecimal encoding and decoding using ModHex alphabet.
See https://developers.yubico.com/OTP/Modhex_Converter.html and https://developers.yubico.com/yubico-c/Manuals/modhex.1.html.

It is a fork of the [`encoding/hex` package](https://golang.org/pkg/encoding/hex/) from the Go standard library.
[`Dump`](https://golang.org/pkg/encoding/hex/#Dump) and [`Dumper`](https://golang.org/pkg/encoding/hex/#Dumper) functions were removed as they are mostly used for debugging, and debugging in ModHex alphabet is not useful.

```sh
go get github.com/AlekSi/modhex@latest
```

```go
import "github.com/AlekSi/modhex"
```
