package yaml_test

import (
	"testing"

	"github.com/nuln/conf"
	"github.com/nuln/conf/conftest"
	cy "github.com/nuln/conf/yaml"
)

func TestYAML(t *testing.T) {
	conftest.Suite(t, cy.New())
}

func TestYAMLRegistration(t *testing.T) {
	available := conf.Available()
	found := false
	for _, name := range available {
		if name == "yaml" {
			found = true
			break
		}
	}
	if !found {
		t.Error("yaml should be registered via init()")
	}
}
