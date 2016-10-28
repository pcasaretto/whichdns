package whichdns

import (
	"net"
	"regexp"
)

type Lookuper interface {
	Lookup(domain string) string
}

type DefaultLookupper struct {
}

var hosts = map[*regexp.Regexp]string{
	regexp.MustCompile(`domaincontrol.com`): "DomainControl",
	regexp.MustCompile(`terra.com.br`):      "Terra",
	regexp.MustCompile(`kinghost.com.br`):   "Kinghost",
}

func (d *DefaultLookupper) Lookup(domain string) string {
	nss, err := net.LookupNS(domain)
	if err != nil {
		return ""
	}
	ns := nss[0]
	for regex, host := range hosts {
		if regex.MatchString(ns.Host) {
			return host
		}
	}
	return ""
}
