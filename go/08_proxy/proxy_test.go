package proxy

import "testing"

// go test -v proxy_test.go proxy.go
func TestProxy(t *testing.T) {
	
	var p Proxy
	
	p = &ProxyClass{}

	p.Do()
}