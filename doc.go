// Package conf provides a unified configuration file reading and writing
// abstraction layer for Go.
//
// It defines a generic [Codec] interface that can be backed by different
// encoding libraries through a codec registration mechanism.
//
// # Supported Codecs
//
//   - json — Go stdlib encoding/json (import _ "github.com/nuln/conf/json")
//   - toml — BurntSushi/toml      (import _ "github.com/nuln/conf/toml")
//   - yaml — gopkg.in/yaml.v3     (import _ "github.com/nuln/conf/yaml")
//
// # Quick Start
//
//	import (
//	    "github.com/nuln/conf"
//	    _ "github.com/nuln/conf/yaml"
//	)
//
//	var cfg AppConfig
//	err := conf.Load("config.yaml", &cfg)
//
// # Import All Codecs
//
//	import _ "github.com/nuln/conf/drivers"
package conf
