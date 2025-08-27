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

func SplitIntoNonBlanks(s, sep string) []string {
	parts := strings.Split(s, sep)
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	parts = FilterInPlace(parts, func(s string) bool { return s != "" })
	return parts
}
