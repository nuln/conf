package conf

import (
	"fmt"
	"sort"
	"sync"
)

var (
	mu     sync.RWMutex
	codecs = make(map[string]Codec)
	// extMap maps file extensions (e.g. ".json") to codec names.
	extMap = make(map[string]string)
)

// Register registers a codec under the given name.
// Adapter sub-packages call this in their init() functions.
// The codec's Extensions() are also indexed for automatic format detection.
func Register(name string, codec Codec) {
	mu.Lock()
	defer mu.Unlock()
	codecs[name] = codec
	for _, ext := range codec.Extensions() {
		extMap[ext] = name
	}
}

// Get returns the codec registered under the given name.
// Returns nil if not found.
func Get(name string) Codec {
	mu.RLock()
	defer mu.RUnlock()
	return codecs[name]
}

// Available returns a sorted list of registered codec names.
func Available() []string {
	mu.RLock()
	defer mu.RUnlock()

	names := make([]string, 0, len(codecs))
	for name := range codecs {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// codecByExt returns the codec registered for the given file extension.
func codecByExt(ext string) (Codec, error) {
	mu.RLock()
	name, ok := extMap[ext]
	mu.RUnlock()

	if !ok {
		return nil, fmt.Errorf("%w: no codec registered for extension %q (available: %v)",
			ErrUnsupportedFormat, ext, Available())
	}

	mu.RLock()
	codec := codecs[name]
	mu.RUnlock()

	return codec, nil
}
