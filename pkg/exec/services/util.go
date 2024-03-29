package services

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"strings"
	"time"
)

const (
	Success = "Success!"
)

var ErrDidNotMatchExpectedOutput = errors.New("did not match Expected Output")

func tcpPortDial(host string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", host, timeout)
	if err != nil {
		return fmt.Errorf("port was not open on a remote host: %w", err)
	}
	_ = conn.Close()
	return nil
}

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
