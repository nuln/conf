package conf

// Codec defines the unified interface for configuration encoding and decoding.
// All adapter implementations must satisfy this interface.
type Codec interface {
	// Encode serializes the given value into bytes.
	// v should be a value (struct, map, etc.) to encode.
	Encode(v any) ([]byte, error)

	// Decode deserializes the given bytes into the value pointed to by v.
	// v must be a pointer to the target type.
	Decode(data []byte, v any) error

	// Extensions returns the file extensions this codec handles,
	// including the leading dot (e.g. [".json"]).
	Extensions() []string
}
