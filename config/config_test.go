//
//
//
package config_test

import(
	"strings"
	"testing"
	"config"
)

var TestProperties *config.Properties 

var TestConfigDataReader = strings.NewReader(`{ "test":true }`)

func TestReadProperties(t *testing.T) {
	p, err := config.ReadProperties(TestConfigDataReader)
	if err != nil {
		t.Fatal(err)
	}
	TestProperties = &p
}

func TestBool(t *testing.T) {
	b, err := TestProperties.Bool("test", "test2")
	if err != nil {
		t.Error(err)
	}

	if b != true {
		t.Error("Expecting true: ", b)
	}
}

