package proxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProxyUrlHttp(t *testing.T) {
	proxy := Proxy{Protocol: HTTP, Host: "example.com"}
	assert.Equal(t, "http://example.com", proxy.Url())
}

func TestProxyUrlSocks5(t *testing.T) {
	proxy := Proxy{Protocol: SOCKS5, Host: "example.com"}
	assert.Equal(t, "socks5://example.com", proxy.Url())
}

func TestProxyUrlLoginPassword(t *testing.T) {
	proxy := Proxy{
		Protocol: HTTP,
		Host:     "example.com",
		User:     "alice",
		Password: "pass",
	}
	assert.Equal(t, "http://alice:pass@example.com", proxy.Url())
}

func TestProxyUrlHttpLoginPasswordPort(t *testing.T) {
	proxy := Proxy{
		Protocol: HTTP,
		User:     "alice",
		Password: "pass",
		Host:     "localhost",
		Port:     1234,
	}
	assert.Equal(t, "http://alice:pass@localhost:1234", proxy.Url())
}
