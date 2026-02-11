package conf_test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/nuln/conf"
	_ "github.com/nuln/conf/json"
	_ "github.com/nuln/conf/toml"
	_ "github.com/nuln/conf/yaml"
)

type appConfig struct {
	Name     string `json:"name"     toml:"name"     yaml:"name"`
	Port     int    `json:"port"     toml:"port"     yaml:"port"`
	Debug    bool   `json:"debug"    toml:"debug"    yaml:"debug"`
	Database struct {
		Host string `json:"host" toml:"host" yaml:"host"`
		Port int    `json:"port" toml:"port" yaml:"port"`
	} `json:"database" toml:"database" yaml:"database"`
}

func sampleConfig() appConfig {
	var cfg appConfig
	cfg.Name = "myapp"
	cfg.Port = 8080
	cfg.Debug = true
	cfg.Database.Host = "localhost"
	cfg.Database.Port = 5432
	return cfg
}

func TestLoadSaveRoundTrip(t *testing.T) {
	formats := []struct {
		ext    string
		format string
	}{
		{".json", "json"},
		{".toml", "toml"},
		{".yaml", "yaml"},
		{".yml", "yaml"},
	}

	for _, f := range formats {
		t.Run(f.format+f.ext, func(t *testing.T) {
			dir := t.TempDir()
			path := filepath.Join(dir, "config"+f.ext)

			original := sampleConfig()

			if err := conf.Save(path, &original); err != nil {
				t.Fatalf("Save failed: %v", err)
			}

			// Verify file was created.
			if _, err := os.Stat(path); err != nil {
				t.Fatalf("file not created: %v", err)
			}

			var loaded appConfig
			if err := conf.Load(path, &loaded); err != nil {
				t.Fatalf("Load failed: %v", err)
			}

			if loaded.Name != original.Name {
				t.Errorf("Name: got %q, want %q", loaded.Name, original.Name)
			}
			if loaded.Port != original.Port {
				t.Errorf("Port: got %d, want %d", loaded.Port, original.Port)
			}
			if loaded.Debug != original.Debug {
				t.Errorf("Debug: got %v, want %v", loaded.Debug, original.Debug)
			}
			if loaded.Database.Host != original.Database.Host {
				t.Errorf("Database.Host: got %q, want %q", loaded.Database.Host, original.Database.Host)
			}
			if loaded.Database.Port != original.Database.Port {
				t.Errorf("Database.Port: got %d, want %d", loaded.Database.Port, original.Database.Port)
			}
		})
	}
}

func TestLoadFromBytesSaveToBytes(t *testing.T) {
	formats := []string{"json", "toml", "yaml"}

	for _, format := range formats {
		t.Run(format, func(t *testing.T) {
			original := sampleConfig()

			data, err := conf.SaveToBytes(&original, format)
			if err != nil {
				t.Fatalf("SaveToBytes failed: %v", err)
			}
			if len(data) == 0 {
				t.Fatal("SaveToBytes returned empty data")
			}

			var loaded appConfig
			if err := conf.LoadFromBytes(data, format, &loaded); err != nil {
				t.Fatalf("LoadFromBytes failed: %v", err)
			}

			if loaded.Name != original.Name {
				t.Errorf("Name: got %q, want %q", loaded.Name, original.Name)
			}
			if loaded.Port != original.Port {
				t.Errorf("Port: got %d, want %d", loaded.Port, original.Port)
			}
		})
	}
}

func TestAvailable(t *testing.T) {
	available := conf.Available()
	expected := []string{"json", "toml", "yaml"}
	for _, name := range expected {
		found := false
		for _, a := range available {
			if a == name {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected %q in Available(), got %v", name, available)
		}
	}
}

func TestLoadUnsupportedFormat(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.xyz")
	if err := os.WriteFile(path, []byte("data"), 0o600); err != nil {
		t.Fatal(err)
	}

	var cfg appConfig
	err := conf.Load(path, &cfg)
	if err == nil {
		t.Fatal("expected error for unsupported format")
	}
	if !errors.Is(err, conf.ErrUnsupportedFormat) {
		t.Errorf("expected ErrUnsupportedFormat, got: %v", err)
	}
}

func TestLoadNoExtension(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config")
	if err := os.WriteFile(path, []byte("data"), 0o600); err != nil {
		t.Fatal(err)
	}

	var cfg appConfig
	err := conf.Load(path, &cfg)
	if err == nil {
		t.Fatal("expected error for file without extension")
	}
	if !errors.Is(err, conf.ErrUnsupportedFormat) {
		t.Errorf("expected ErrUnsupportedFormat, got: %v", err)
	}
}

func TestLoadNonExistentFile(t *testing.T) {
	var cfg appConfig
	err := conf.Load("/nonexistent/config.json", &cfg)
	if err == nil {
		t.Fatal("expected error for nonexistent file")
	}
}

func TestLoadFromBytesUnknownFormat(t *testing.T) {
	var cfg appConfig
	err := conf.LoadFromBytes([]byte("data"), "unknown", &cfg)
	if err == nil {
		t.Fatal("expected error for unknown format")
	}
	if !errors.Is(err, conf.ErrUnsupportedFormat) {
		t.Errorf("expected ErrUnsupportedFormat, got: %v", err)
	}
}

func TestSaveToBytesUnknownFormat(t *testing.T) {
	cfg := sampleConfig()
	_, err := conf.SaveToBytes(&cfg, "unknown")
	if err == nil {
		t.Fatal("expected error for unknown format")
	}
	if !errors.Is(err, conf.ErrUnsupportedFormat) {
		t.Errorf("expected ErrUnsupportedFormat, got: %v", err)
	}
}

func TestGet(t *testing.T) {
	codec := conf.Get("json")
	if codec == nil {
		t.Error("Get(\"json\") returned nil")
	}

	codec = conf.Get("nonexistent")
	if codec != nil {
		t.Error("Get(\"nonexistent\") should return nil")
	}
}
