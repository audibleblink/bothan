package modules

import "net/http"

// Registry holds the list OSTs for the app
var Registry = make(map[string]ToolInt)

// ToolInt is the interface which all OST modules should satisfy
type ToolInt interface {
	IsBad(*Query) bool
	Name() string
	SetName(string)
}

// Tool is the structure by which individual OSTs will be instantiated
type Tool struct{ name string }

func (t *Tool) Name() string {
	return t.name
}

func (t *Tool) SetName(name string) {
	t.name = name
}

// Query is a wrapper around raw HTTP requests
type Query struct {
	Request     *http.Request
	RawResponse []byte
}

// MasscanRecord holds the structure for masscan records that get read
// in from STDIN when using masscan with the -oD - output flag
type MasscanRecord struct {
	IP        string `json:"ip"`
	Timestamp string `json:"timestamp"`
	Port      int    `json:"port"`
}
