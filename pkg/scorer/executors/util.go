package executors

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net"
	"net/url"
	"strings"
	"time"
)

const (
	Success = "Success!"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

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

// IPToStringSlice converts a slice of net.IP to a slice of strings.
func IPToStringSlice(ips []net.IP) []string {
	stringSlice := make([]string, len(ips))
	for i, ip := range ips {
		stringSlice[i] = ip.String()
	}
	return stringSlice
}

//func UpdateScorerProperties(ss ScoreService, properties map[string]string) (err error) {
//	//defer func() {
//	//	if x := recover(); x != nil {
//	//		switch x := x.(type) {
//	//		case string:
//	//			err = fmt.Errorf("%w: %s", ErrPanic, x)
//	//		case error:
//	//			err = x
//	//		default:
//	//			err = ErrUnknownPanic
//	//		}
//	//	}
//	//}()
//
//	rv := reflect.ValueOf(ss).Elem()
//	for key, val := range properties {
//		if val != "" { // Eliminate unnecessary default value
//			rf := rv.FieldByName(key)
//			rf.SetString(val)
//		}
//	}
//	return nil
//}

//func UnmarshallProperties(properties any, propertyType interface{}) (err error) {
