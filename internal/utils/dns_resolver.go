package utils

import (
	"net"
	"time"
	"context"
)

type DNSResolver struct {
	Resolver *net.Resolver
}

func NewDNSResolver() *DNSResolver {
	return &DNSResolver{
		Resolver: &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{
					Timeout: time.Second * 5,
				}
				return d.DialContext(ctx, network, address)
			},
		},
	}
}

func (dr *DNSResolver) LookupHost(host string) ([]string, error) {
	return dr.Resolver.LookupHost(context.Background(), host)
}
