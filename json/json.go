// Package json provides a JSON codec for the conf package.
// Import this package to register the "json" codec:
//
//	import _ "github.com/nuln/conf/json"
package json

import (
	"encoding/json"

	"github.com/nuln/conf"
)

func init() {
	conf.Register("json", New())
}

// New returns a new JSON codec.
func New() conf.Codec {
	return &jsonCodec{}
}

type jsonCodec struct{}

func (c *jsonCodec) Encode(v any) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

func (c *jsonCodec) Decode(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func (c *jsonCodec) Extensions() []string {
	return []string{".json"}
}

var _ conf.Codec = (*jsonCodec)(nil)
