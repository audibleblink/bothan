package modules

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

const (
	sig1 = "c385c13e5100a836c6c05a03ad82be1bb9fee7245478c47d88d771f57cd12581"
)

// CobaltStrike type
type CobaltStrike struct {
	Tool
}

func init() {
	tool := &CobaltStrike{}
	tool.SetName("CobaltStrike")
	Registry["cobaltstrike"] = tool
}

// IsBad evaluates confidence in a target's badness
func (t CobaltStrike) IsBad(req *Query) bool {
	headers := strings.Split(string(req.RawResponse), "\r\n")

	if spaceSig(headers) {
		Registry["cobaltstrike"].SetName("CobaltStrike < v3.13")
		return true
	}

	return false
}

func spaceSig(headers []string) bool {
	subHeaders := headers[0:2]
	subHeaders = append(subHeaders, headers[3]+"\r\n\r\n")
	joined := strings.Join(subHeaders, "\r\n")

	sha := sha256.New()
	sha.Write([]byte(joined))

	hash := fmt.Sprintf("%x", sha.Sum(nil))
	return hash == sig1
}
