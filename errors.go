package conf

import "errors"

// Sentinel errors returned by conf functions.
var (
	// ErrUnsupportedFormat is returned when no codec is registered for the
	// requested format or file extension.
	ErrUnsupportedFormat = errors.New("conf: unsupported format")
)
