// Package drivers is a convenience package that registers all built-in
// codecs. Import it with a blank identifier to make all codecs available:
//
//	import _ "github.com/nuln/conf/drivers"
package drivers

import (
	"github.com/nuln/conf"
	_ "github.com/nuln/conf/json"
	_ "github.com/nuln/conf/toml"
	_ "github.com/nuln/conf/yaml"
)

// Init ensures all built-in codecs are registered.
// This is called automatically by importing the package.
func Init() {}

// List returns a list of all registered codecs.
func List() []string {
	return conf.Available()
}

// Load reads the file at path, detects the format from the file extension,
// and decodes its contents into v. v must be a pointer.
//
// This is a convenience wrapper for conf.Load.
func Load(path string, v any) error {
	return conf.Load(path, v)
}

// Save encodes v and writes the result to the file at path.
// The format is detected from the file extension.
// Parent directories are created automatically if they do not exist.
//
// This is a convenience wrapper for conf.Save.
func Save(path string, v any) error {
	return conf.Save(path, v)
}
