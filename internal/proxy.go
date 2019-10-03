package proxy

import "fmt"

type Protocol string

const (
	HTTP   = "http"
	SOCKS5 = "socks5"
)

type Proxy struct {
	Protocol Protocol
	User     string
	Password string
	Host     string
	Port     uint16
}

func (p Proxy) Url() string {
	protocol := p.Protocol
	auth := ""
	host := p.Host
	port := ""

	if p.User != "" {
		auth = fmt.Sprintf("%v:%v@", p.User, p.Password)
	}
	if p.Port != 0 {
		port = fmt.Sprintf(":%v", p.Port)
	}

	return fmt.Sprintf("%v://%v%v%v", protocol, auth, host, port)
}
