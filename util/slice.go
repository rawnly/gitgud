package util

import "strings"

// RemoveEmptyStrings removes empty strings from an array of strings
func RemoveEmptyStrings(arr []string) []string {
	var items []string

	for _, item := range arr {
		trimmed := strings.TrimRight(strings.TrimSpace(item), "\n")

		if len(trimmed) > 0 {
			items = append(items, trimmed)
		}
	}

	return items
}
