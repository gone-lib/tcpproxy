# tcpproxy

Package tcpproxy lets users build TCP proxies, optionally making routing decisions based on HTTP/1 Host headers and the SNI hostname in TLS connections.

Calling Run (or Start) on a proxy also starts all the necessary listeners.

For each accepted connection, the rules for that ipPort are matched, in order. If one matches (currently HTTP Host, SNI, or always), then the connection is handed to the target.

The two predefined Target implementations are:

- DialProxy, proxying to another address (use the To func to return a DialProxy value),

- TargetListener, making the matched connection available via a net.Listener.Accept call.

But Target is an interface, so you can also write your own.

Note that tcpproxy does not do any TLS encryption or decryption. It only (via DialProxy) copies bytes around. The SNI hostname in the TLS header is unencrypted, for better or worse.

This package makes no API stability promises. If you depend on it, vendor it.

## Install

Install TLSRouter via `go get`:

```shell
go get github.com/gone-lib/tcpproxy
```

## Usage

```go
package main

import (
	"log"

	"github.com/gone-lib/tcpproxy"
)

func main() {
	var p tcpproxy.Proxy
	// p.AddHTTPHostRoute(":80", "foo.com", tcpproxy.To("10.0.0.1:8081"))
	// p.AddHTTPHostRoute(":80", "bar.com", tcpproxy.To("10.0.0.2:8082"))
	// p.AddRoute(":80", tcpproxy.To("10.0.0.1:8081")) // fallback
	// p.AddSNIRoute(":443", "foo.com", tcpproxy.To("10.0.0.1:4431"))
	// p.AddSNIRoute(":443", "bar.com", tcpproxy.To("10.0.0.2:4432"))
	// p.AddRoute(":443", tcpproxy.To("10.0.0.1:4431")) // fallback
	p.AddRoute(":9999", tcpproxy.To("127.0.0.1:8888")) // fallback
	log.Fatal(p.Run())
}
```
