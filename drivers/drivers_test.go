package drivers_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/nuln/conf/drivers"
)

type testConfig struct {
	Name string `json:"name" yaml:"name" toml:"name"`
	Port int    `json:"port" yaml:"port" toml:"port"`
}

func TestDriversLoadSave(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.json")

	original := testConfig{Name: "drivers-test", Port: 1234}

	// Test Save
	if err := drivers.Save(path, &original); err != nil {
		t.Fatalf("drivers.Save failed: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("file not created: %v", err)
	}

	// Test Load
	var loaded testConfig
	if err := drivers.Load(path, &loaded); err != nil {
		t.Fatalf("drivers.Load failed: %v", err)
	}

	if loaded.Name != original.Name || loaded.Port != original.Port {
		t.Errorf("loaded data mismatch: got %+v, want %+v", loaded, original)
	}
}

func TestDriversLoadSaveAll(t *testing.T) {
	dir := t.TempDir()

	formats := []struct {
		name string
		ext  string
	}{
		{"json", ".json"},
		{"toml", ".toml"},
		{"yaml", ".yaml"},
		{"yml", ".yml"},
	}

	for _, f := range formats {
		t.Run(f.name, func(t *testing.T) {
			path := filepath.Join(dir, "config"+f.ext)
			original := testConfig{Name: "test-" + f.name, Port: 8888}

			// drivers.Save detects format from extension
			if err := drivers.Save(path, &original); err != nil {
				t.Fatalf("drivers.Save (%s) failed: %v", f.ext, err)
			}

			// drivers.Load also detects format
			var loaded testConfig
			if err := drivers.Load(path, &loaded); err != nil {
				t.Fatalf("drivers.Load (%s) failed: %v", f.ext, err)
			}

			if loaded.Name != original.Name {
				t.Errorf("format %s: expected Name %q, got %q", f.ext, original.Name, loaded.Name)
			}
		})
	}
}

func TestList(t *testing.T) {
	list := drivers.List()
	if len(list) < 3 {
		t.Errorf("expected at least 3 codecs, got %v", list)
	}
}
