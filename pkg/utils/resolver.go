package utils

import (
    "fmt"
    "github.com/miekg/dns"
)

// ResolveIPAdvanced resolves the IP addresses for the given host using advanced DNS resolution
func ResolveIPAdvanced(host string) ([]string, error) {
    client := new(dns.Client)
    message := new(dns.Msg)
    message.SetQuestion(dns.Fqdn(host), dns.TypeA)
    message.RecursionDesired = true

    // Use Google's public DNS server
    dnsServer := "8.8.8.8:53"

    response, _, err := client.Exchange(message, dnsServer)
    if err != nil {
        return nil, fmt.Errorf("failed to resolve IP: %v", err)
    }

    if response.Rcode != dns.RcodeSuccess {
        return nil, fmt.Errorf("failed to get a valid answer: %v", response.Rcode)
    }

    var ipStrings []string
    for _, answer := range response.Answer {
        if aRecord, ok := answer.(*dns.A); ok {
            ipStrings = append(ipStrings, aRecord.A.String())
        }
    }

    return ipStrings, nil
}