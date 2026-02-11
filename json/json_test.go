package json_test

import (
	"testing"

	"github.com/nuln/conf"
	"github.com/nuln/conf/conftest"
	cj "github.com/nuln/conf/json"
)

func TestJSON(t *testing.T) {
	conftest.Suite(t, cj.New())
}

func TestJSONRegistration(t *testing.T) {
	available := conf.Available()
	found := false
	for _, name := range available {
		if name == "json" {
			found = true
			break
		}
	}
	if !found {
		t.Error("json should be registered via init()")
	}
}
