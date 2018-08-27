// api-pub-net/models
//
// This package contains all models used in this service.

package models

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