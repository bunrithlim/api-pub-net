// api-pub-net/models
//
// This package contains all models used in this service.

package models

import (
	"net"
)

// IPAddress is a struct we use to represent JSON API responses for IP addresses.
type IPAddress struct {
	IP string `json:"ip"`
}

// TimeUTC is a struct we use to represent JSON API responses for UTC Times.
type TimeUTC struct {
	UTC string `json:"utc"`
}

// RequestInfo is a struct we use to represent JSON API responses about the request.
type RequestInfo struct {
	IP string `json:"ip"`
	UserAgent string `json:"user-agent"`
	Referrer string `json:"referer"`
}

// RequestHeader is a struct we use to represent JSON API responses about the request.
type RequestHeader struct {
	Header string `json:"header"`
	Value string `json:"value"`
}

// RequestHeaders is a struct we use to represent JSON API responses about the request.
type RequestHeaders struct {
	Method string `json:"method"`
	URL string `json:"url"`
	Protocol string `json:"protocol"`
	Host string `json:"host"`
	RemoteIpAddress string `json:"remote_address"`
	Headers []RequestHeader `json:"headers"`
}

// DNSHost is a struct we use to represent JSON API responses for DNS Hosts.
type DNSAll struct {
	Cname string `json:"cname,omitempty"`
	Host_addrs []string `json:"host,omitempty"`
	Mx []*net.MX `json:"mx,omitempty"`
	Ns []*net.NS `json:"ns,omitempty"`
	Txts []string `json:"txts,omitempty"`
	IP []net.IP `json:"ip,omitempty"`
}



