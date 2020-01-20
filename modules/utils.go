package modules

import (
	"strings"
)

func headerKeys(headers []string) (cleaned []string) {
	for _, el := range headers {
		pair := strings.Split(el, ":")
		if pair[0] == "" {
			continue
		}
		cleaned = append(cleaned, pair[0])
	}
	return
}

func equalSlice(a, b []string) (ok bool) {
	if len(a) != len(b) {
		return
	}

	for i, el := range a {
		if el != b[i] {
			return
		}
	}
	return true
}
