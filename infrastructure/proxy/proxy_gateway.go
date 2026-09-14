package proxy

import (
	"math/rand"
)

type ProxyGateway struct {
	ProxyPool []string
}

func NewProxyGateway(proxies []string) *ProxyGateway {
	return &ProxyGateway{ProxyPool: proxies}
}

func (p *ProxyGateway) GetNextIP() string {
	if len(p.ProxyPool) == 0 {
		return "direct"
	}
	return p.ProxyPool[rand.Intn(len(p.ProxyPool))]
}
