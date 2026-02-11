// Package toml provides a TOML codec for the conf package.
// Import this package to register the "toml" codec:
//
//	import _ "github.com/nuln/conf/toml"
package toml

import (
	"bytes"

	"github.com/BurntSushi/toml"

	"github.com/nuln/conf"
)

func init() {
	conf.Register("toml", New())
}

// New returns a new TOML codec.
func New() conf.Codec {
	return &tomlCodec{}
}

type tomlCodec struct{}

func (c *tomlCodec) Encode(v any) ([]byte, error) {
	var buf bytes.Buffer
	encoder := toml.NewEncoder(&buf)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *tomlCodec) Decode(data []byte, v any) error {
	_, err := toml.Decode(string(data), v)
	return err
}

func (c *tomlCodec) Extensions() []string {
	return []string{".toml"}
}

var _ conf.Codec = (*tomlCodec)(nil)
