package util

import "strings"

func ContainsIgnoringCase(string, fragment string) bool {
	string = strings.ToLower(string)
	fragment = strings.ToLower(fragment)
	return strings.Contains(string, fragment)
}

func TakeStr(s string, n int) string {
	if len(s) < n {
		return s
	}
	return s[:n]
}
