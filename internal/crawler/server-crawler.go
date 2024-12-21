package crawler

import (
	"distributed-crawler/internal/models"
	"distributed-crawler/internal/storage"
	"distributed-crawler/internal/utils"
	"log"
	"net"
)

type Crawler struct {
	Storage     storage.Storage
	DNSResolver *utils.DNSResolver
}

func NewCrawler(storage storage.Storage) *Crawler {
	return &Crawler{
		Storage:     storage,
		DNSResolver: utils.NewDNSResolver(),
	}
}

func (c *Crawler) Crawl(startURL string) {
	c.Storage.Save(models.URL{Address: startURL, Priority: 0})
	CrawlURL(startURL, c.handleLink)
}

func (c *Crawler) handleLink(link string) {
	hosts, err := c.DNSResolver.LookupHost(link)
	if err != nil {
		log.Printf("DNS lookup failed for %s: %v", link, err)
		return
	}
	for _, host := range hosts {
		c.resolveAndSaveHost(host)
	}
	c.Storage.Save(models.URL{Address: link, Priority: 1})
}

func (c *Crawler) resolveAndSaveHost(host string) {
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Printf("IP lookup failed for %s: %v", host, err)
		return
	}
	for _, ip := range ips {
		c.reverseLookupAndSaveIP(ip)
	}
}

func (c *Crawler) reverseLookupAndSaveIP(ip net.IP) {
	names, err := net.LookupAddr(ip.String())
	if err != nil {
		log.Printf("Reverse DNS lookup failed for %s: %v", ip, err)
		return
	}
	for _, name := range names {
		log.Printf("Resolved name for IP %s: %s", ip, name)
		c.Storage.Save(models.URL{Address: name, Priority: 1})
	}
}
