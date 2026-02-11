// Package yaml provides a YAML codec for the conf package.
// Import this package to register the "yaml" codec:
//
//	import _ "github.com/nuln/conf/yaml"
package yaml

import (
	"gopkg.in/yaml.v3"

	"github.com/nuln/conf"
)

func init() {
	conf.Register("yaml", New())
}

// New returns a new YAML codec.
func New() conf.Codec {
	return &yamlCodec{}
}

type yamlCodec struct{}

func (c *yamlCodec) Encode(v any) ([]byte, error) {
	return yaml.Marshal(v)
}

func (c *yamlCodec) Decode(data []byte, v any) error {
	return yaml.Unmarshal(data, v)
}

func (c *yamlCodec) Extensions() []string {
	return []string{".yaml", ".yml"}
}

var _ conf.Codec = (*yamlCodec)(nil)
