package toml_test

import (
	"testing"

	"github.com/nuln/conf"
	"github.com/nuln/conf/conftest"
	ct "github.com/nuln/conf/toml"
)

func TestTOML(t *testing.T) {
	conftest.Suite(t, ct.New())
}

func TestTOMLRegistration(t *testing.T) {
	available := conf.Available()
	found := false
	for _, name := range available {
		if name == "toml" {
			found = true
			break
		}
	}
	if !found {
		t.Error("toml should be registered via init()")
	}
}
