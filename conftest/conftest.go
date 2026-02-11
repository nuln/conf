// Package conftest provides a conformance test suite for conf.Codec
// implementations. Adapter packages should use Suite in their tests
// to verify correct behavior.
package conftest

import (
	"reflect"
	"testing"

	"github.com/nuln/conf"
)

// testConfig is the canonical struct used across all conformance tests.
type testConfig struct {
	Name     string            `json:"name"     toml:"name"     yaml:"name"`
	Port     int               `json:"port"     toml:"port"     yaml:"port"`
	Debug    bool              `json:"debug"    toml:"debug"    yaml:"debug"`
	Rate     float64           `json:"rate"     toml:"rate"     yaml:"rate"`
	Tags     []string          `json:"tags"     toml:"tags"     yaml:"tags"`
	Metadata map[string]string `json:"metadata" toml:"metadata" yaml:"metadata"`
	Database dbConfig          `json:"database" toml:"database" yaml:"database"`
}

type dbConfig struct {
	Host string `json:"host" toml:"host" yaml:"host"`
	Port int    `json:"port" toml:"port" yaml:"port"`
}

func sampleConfig() testConfig {
	return testConfig{
		Name:  "myapp",
		Port:  8080,
		Debug: true,
		Rate:  3.14,
		Tags:  []string{"web", "api"},
		Metadata: map[string]string{
			"version": "1.0",
			"env":     "dev",
		},
		Database: dbConfig{
			Host: "localhost",
			Port: 5432,
		},
	}
}

// Suite runs a conformance test suite against a codec implementation.
//
//nolint:gocyclo // complexity is from sequential subtests, not logic branching
func Suite(t *testing.T, codec conf.Codec) {
	t.Helper()

	t.Run("RoundTrip", func(t *testing.T) {
		original := sampleConfig()

		data, err := codec.Encode(original)
		if err != nil {
			t.Fatalf("Encode failed: %v", err)
		}
		if len(data) == 0 {
			t.Fatal("Encode returned empty data")
		}

		var decoded testConfig
		if err := codec.Decode(data, &decoded); err != nil {
			t.Fatalf("Decode failed: %v", err)
		}

		if !reflect.DeepEqual(original, decoded) {
			t.Errorf("round-trip mismatch:\n  original: %+v\n  decoded:  %+v", original, decoded)
		}
	})

	t.Run("NestedStruct", func(t *testing.T) {
		original := sampleConfig()

		data, err := codec.Encode(original)
		if err != nil {
			t.Fatalf("Encode failed: %v", err)
		}

		var decoded testConfig
		if err := codec.Decode(data, &decoded); err != nil {
			t.Fatalf("Decode failed: %v", err)
		}

		if decoded.Database.Host != "localhost" {
			t.Errorf("expected nested host %q, got %q", "localhost", decoded.Database.Host)
		}
		if decoded.Database.Port != 5432 {
			t.Errorf("expected nested port %d, got %d", 5432, decoded.Database.Port)
		}
	})

	t.Run("ZeroValue", func(t *testing.T) {
		original := testConfig{}

		data, err := codec.Encode(original)
		if err != nil {
			t.Fatalf("Encode failed: %v", err)
		}

		var decoded testConfig
		if err := codec.Decode(data, &decoded); err != nil {
			t.Fatalf("Decode failed: %v", err)
		}

		if decoded.Name != "" {
			t.Errorf("expected empty name, got %q", decoded.Name)
		}
		if decoded.Port != 0 {
			t.Errorf("expected zero port, got %d", decoded.Port)
		}
	})

	t.Run("Extensions", func(t *testing.T) {
		exts := codec.Extensions()
		if len(exts) == 0 {
			t.Error("Extensions() returned empty slice")
		}
		for _, ext := range exts {
			if ext == "" || ext[0] != '.' {
				t.Errorf("extension should start with dot, got %q", ext)
			}
		}
	})

	t.Run("DecodeInvalidData", func(t *testing.T) {
		var decoded testConfig
		err := codec.Decode([]byte("<<<invalid>>>"), &decoded)
		if err == nil {
			t.Error("expected error when decoding invalid data, got nil")
		}
	})
}
