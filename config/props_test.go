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

var TestConfigDataReader = strings.NewReader(`{
	"bool1":true,
	"bool2":false,
	"level2":{
		"bool3":true,
		"bool4":false,
		"bool5":[ true, false ]
	}
}`)

func TestReadProperties(t *testing.T) {
	props, err := config.ReadProperties(TestConfigDataReader)
	if err == nil {
		t.Log("Success reading config properties from JSON string.")
	} else {
		t.Fatal("Error reading config properties from JSON string:", err)
	}
	TestProperties = props
}

func TestBool1(t *testing.T) {
	if(TestProperties == nil) {
		t.FailNow()
	}

	b, err := TestProperties.Bool("bool1")
	if err == nil {
		t.Log("Success retrieving Bool value from property 'bool1'.")
	} else {
		t.Fatal("Error retrieving Bool value from property 'bool1':", err)
	}

	if b == true {
		t.Log("Bool value for 'bool1' is true.")
	} else {
		t.Fatal("Bool value for 'bool1' is not true.")
	}
}

func TestBool2(t *testing.T) {
	if(TestProperties == nil) {
		t.FailNow()
	}

	b, err := TestProperties.Bool("bool2")
	if err == nil {
		t.Log("Success retrieving Bool value from property 'bool2'.")
	} else {
		t.Fatal("Error retrieving Bool value from property 'bool2':", err)
	}

	if b == false {
		t.Log("Bool value for 'bool2' is false.")
	} else {
		t.Fatal("Bool value for 'bool2' is not false.")
	}
}

func TestBool3(t *testing.T) {
	if(TestProperties == nil) {
		t.FailNow()
	}

	b, err := TestProperties.Bool("level2.bool3")
	if err == nil {
		t.Log("Success retrieving Bool value from property 'bool3'.")
	} else {
		t.Fatal("Error retrieving Bool value from property 'bool3':", err)
	}

	if b == true {
		t.Log("Bool value for 'bool3' is true.")
	} else {
		t.Fatal("Bool value for 'bool3' is not true.")
	}
}

func TestBool4(t *testing.T) {
	if(TestProperties == nil) {
		t.FailNow()
	}

	b, err := TestProperties.Bool("level2", "bool4")
	if err == nil {
		t.Log("Success retrieving Bool value from property 'bool4'.")
	} else {
		t.Fatal("Error retrieving Bool value from property 'bool4':", err)
	}

	if b == false {
		t.Log("Bool value for 'bool4' is false.")
	} else {
		t.Fatal("Bool value for 'bool4' is not false.")
	}
}

func TestBool5(t *testing.T) {
	if(TestProperties == nil) {
		t.FailNow()
	}

	b, err := TestProperties.Bool("level2.bool5[0]")
	if err == nil {
		t.Log("Success retrieving Bool value from property 'bool5[0]'.")
	} else {
		t.Fatal("Error retrieving Bool value from property 'bool5[0]':", err)
	}

	if b == true {
		t.Log("Bool value for 'bool5[0]' is true.")
	} else {
		t.Fatal("Bool value for 'bool5[0]' is not true.")
	}
	
	b, err = TestProperties.Bool("level2", "bool5", 1)
	if err == nil {
		t.Log("Success retrieving Bool value from property 'bool5[1]'.")
	} else {
		t.Fatal("Error retrieving Bool value from property 'bool5[1]':", err)
	}

	if b == false {
		t.Log("Bool value for 'bool5[1]' is false.")
	} else {
		t.Fatal("Bool value for 'bool5[1]' is not false.")
	}
}