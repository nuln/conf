# conf

[![CI](https://github.com/nuln/conf/actions/workflows/ci.yml/badge.svg)](https://github.com/nuln/conf/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/nuln/conf.svg)](https://pkg.go.dev/github.com/nuln/conf)

A unified Go configuration file library with pluggable codecs. Read and write config files in multiple formats through a single API.

## Features

- **Unified Interface**: Read and write configuration files via a single `Codec` interface.
- **Pluggable Codecs**: Built-in support for `json`, `toml`, and `yaml`.
- **Auto Detection**: Format automatically detected from file extension.
- **Easy Registration**: Codecs register themselves via `init()` — just import and use.
- **Thread-Safe**: Registry is safe for concurrent use.
- **Extensible**: Add new formats by implementing the `Codec` interface.

## Supported Codecs

| Format | Import | Registration Name | Extensions |
|--------|--------|-------------------|------------|
| JSON | `github.com/nuln/conf/json` | `"json"` | `.json` |
| TOML | `github.com/nuln/conf/toml` | `"toml"` | `.toml` |
| YAML | `github.com/nuln/conf/yaml` | `"yaml"` | `.yaml`, `.yml` |

## Installation

```bash
go get github.com/nuln/conf
```

## Quick Start

### One-Stop Usage (Recommended)

Import the `drivers` package to automatically register all codecs and use the convenience API:

```go
import "github.com/nuln/conf/drivers"

type AppConfig struct {
    Host string `yaml:"host" json:"host"`
    Port int    `yaml:"port" json:"port"`
}

func main() {
    var cfg AppConfig
    // Load detects format from extension (.yaml)
    if err := drivers.Load("config.yaml", &cfg); err != nil {
        panic(err)
    }
    
    // Save also detects format
    cfg.Port = 8080
    if err := drivers.Save("config.yaml", &cfg); err != nil {
        panic(err)
    }
}
```

### Manual Registration

If you only need a specific codec and want to minimize dependencies:

```go
import (
    "github.com/nuln/conf"
    _ "github.com/nuln/conf/yaml" // auto-register yaml codec
)

var cfg AppConfig
err := conf.Load("config.yaml", &cfg)
```

### 4. Switch Formats

To switch formats, just change the import and file extension — no other code changes required:

```go
import _ "github.com/nuln/conf/json"   // Switch to JSON
import _ "github.com/nuln/conf/toml"   // Switch to TOML
```

## Advanced Usage

### Bytes API

Work with configuration data in memory without files:

```go
// Encode to bytes
data, err := conf.SaveToBytes(&cfg, "json")

// Decode from bytes
var cfg2 AppConfig
err = conf.LoadFromBytes(data, "json", &cfg2)
```

### Direct Codec Access

```go
codec := conf.Get("json")
data, _ := codec.Encode(cfg)
```

### List Registered Codecs

```go
fmt.Println(conf.Available()) // [json toml yaml]
```

## Development

The project includes a `Makefile` for standard development tasks:

```bash
make all      # Run fmt, tidy, lint and test
make test     # Run tests with race detection
make lint     # Run golangci-lint
make build    # Build the project
make coverage # Generate coverage report
make clean    # Clean artifacts and test cache
make help     # Show all targets
```

## Contributing

New codecs can be added by implementing the `conf.Codec` interface and registering them via `conf.Register`. See [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

Apache License 2.0
