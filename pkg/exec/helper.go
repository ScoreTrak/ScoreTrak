package exec

import "strings"

func IsSecure(s string) bool {
	if containsString([]string{"https", "tls", "ssl", "ldaps"}, strings.ToLower(s)) {
		return true
	}
	return false
}

func containsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
