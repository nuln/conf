package conf

import (
	"fmt"
	"os"
	"path/filepath"
)

// Load reads the file at path, detects the format from the file extension,
// and decodes its contents into v. v must be a pointer.
func Load(path string, v any) error {
	codec, err := codecForPath(path)
	if err != nil {
		return err
	}

	data, err := os.ReadFile(path) //nolint:gosec // path is user-provided by design
	if err != nil {
		return fmt.Errorf("conf: reading %s: %w", path, err)
	}

	if err := codec.Decode(data, v); err != nil {
		return fmt.Errorf("conf: decoding %s: %w", path, err)
	}
	return nil
}

// Save encodes v and writes the result to the file at path.
// The format is detected from the file extension.
// Parent directories are created automatically if they do not exist.
func Save(path string, v any) error {
	codec, err := codecForPath(path)
	if err != nil {
		return err
	}

	data, err := codec.Encode(v)
	if err != nil {
		return fmt.Errorf("conf: encoding %s: %w", path, err)
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o750); err != nil {
		return fmt.Errorf("conf: creating directory %s: %w", dir, err)
	}

	if err := os.WriteFile(path, data, 0o600); err != nil {
		return fmt.Errorf("conf: writing %s: %w", path, err)
	}
	return nil
}

// LoadFromBytes decodes data in the named format into v.
// format is a registered codec name (e.g. "json", "yaml", "toml").
func LoadFromBytes(data []byte, format string, v any) error {
	codec := Get(format)
	if codec == nil {
		return fmt.Errorf("%w: %q (available: %v)", ErrUnsupportedFormat, format, Available())
	}
	if err := codec.Decode(data, v); err != nil {
		return fmt.Errorf("conf: decoding %s: %w", format, err)
	}
	return nil
}

// SaveToBytes encodes v using the named format and returns the bytes.
// format is a registered codec name (e.g. "json", "yaml", "toml").
func SaveToBytes(v any, format string) ([]byte, error) {
	codec := Get(format)
	if codec == nil {
		return nil, fmt.Errorf("%w: %q (available: %v)", ErrUnsupportedFormat, format, Available())
	}
	data, err := codec.Encode(v)
	if err != nil {
		return nil, fmt.Errorf("conf: encoding %s: %w", format, err)
	}
	return data, nil
}

// codecForPath returns the codec matched by the file extension of path.
func codecForPath(path string) (Codec, error) {
	ext := filepath.Ext(path)
	if ext == "" {
		return nil, fmt.Errorf("%w: file %q has no extension", ErrUnsupportedFormat, path)
	}
	return codecByExt(ext)
}
