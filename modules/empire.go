package modules

import (
	"strings"
)

// Empire type
type Empire Tool

func init() {
	Registry["empire"] = Empire{}
}

// IsBad evaluates confidence in a target's badness
func (t Empire) IsBad(req *Query) bool {
	headers := strings.Split(string(req.RawResponse), "\r\n")
	keys := headerKeys(headers)
	return empireSig1(headers) && empireHead(keys) && empireSig2(headers)
}

func empireSig1(headers []string) bool {
	bytes := []byte{0x2e, 0x32, 0x32, 0x36, 0x49, 0x57, 0x48, 0x56, 0x46, 0x54, 0x56, 0x56, 0x46, 0x29, 0x2d}
	var out []byte
	for _, v := range bytes {
		out = append(out, v^0x66)
	}
	return string(out) == headers[0]
}

func empireSig2(headers []string) (ok bool) {
	bytes := []byte{0x25, 0x07, 0x05, 0x0e, 0x03, 0x4b, 0x25, 0x09, 0x08, 0x12, 0x14, 0x09, 0x0a, 0x5c, 0x46, 0x08, 0x09, 0x4b, 0x05, 0x07, 0x05, 0x0e, 0x03, 0x4a, 0x46, 0x08, 0x09, 0x4b, 0x15, 0x12, 0x09, 0x14, 0x03, 0x4a, 0x46, 0x0b, 0x13, 0x15, 0x12, 0x4b, 0x14, 0x03, 0x10, 0x07, 0x0a, 0x0f, 0x02, 0x07, 0x12, 0x03}
	var out []byte
	for _, v := range bytes {
		out = append(out, v^0x66)
	}
	return string(out) == headers[3]
}

func empireHead(headers []string) (ok bool) {
	wantKeys := []string{
		"Connection",
		"Content-Length",
		"Cache-Control",
		"Content-Type",
		"Date",
		"Expires",
		"Pragma",
		"Server",
	}

	alsoWant := []string{
		"Connection",
		"Content-Length",
		"Cache-Control",
		"Content-Type",
	}

	ok = equalSlice(headers[1:], wantKeys) || equalSlice(headers[1:5], alsoWant)
	return
}
