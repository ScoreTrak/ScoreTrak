package exec

import "strings"

func IsSecure(s string) bool {
	if ContainsString([]string{"https", "tls", "ssl", "ldaps"}, strings.ToLower(s)) {
		return true
	}
	return false
}

func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
