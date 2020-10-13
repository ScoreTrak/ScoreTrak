package exec

import "strings"

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
