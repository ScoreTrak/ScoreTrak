package services

import (
	"errors"
	"net/url"
	"strings"
)

const (
	Success = "Success!"
)

var ErrDidNotMatchExpectedOutput = errors.New("did not match Expected Output")

func IsSecure(s string) bool {
	return ContainsString([]string{"https", "tls", "ssl", "ldaps"}, strings.ToLower(s))
}

func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ConstructURI(port, subdomain, host, path, scheme string) *url.URL {
	if port == "" {
		if strings.ToLower(scheme) == "https" {
			port = "443"
		} else {
			port = "80"
		}
	}
	if subdomain != "" && subdomain[len(subdomain)-1:] != "." {
		subdomain += "."
	}
	return &url.URL{Path: path, Scheme: scheme, Host: subdomain + host + ":" + port}
}
