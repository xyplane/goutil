//
//
//
package config_test

import(
	"config"
	"strings"
	"testing"
)

var TestBoolConfigData = `{
	"bool1":true,
	"bool2":false,
	"level2":{
		"bool3":true,
		"bool4":false,
		"bool5":[ true, false ]
	}
}`

func TestBool(t *testing.T) {

	t.Log("Read the following JSON config data:\n" + TestBoolConfigData)

	properties, err := config.ReadProperties(strings.NewReader(TestBoolConfigData))
	if err == nil {
		t.Log("Success reading config properties.")
	} else {
		t.Fatal("Error reading config properties:", err)
	}
	
	var b bool

	b, err = properties.Bool("bool1")
	if err == nil {
		if b == true {
			t.Log("Bool value for 'bool1' is true.")
		} else {
			t.Error("Bool value for 'bool1' is not true.")
		}
	} else {
		t.Error("Error getting bool value from property 'bool1':", err)
	}


	b, err = properties.Bool("bool2")
	if err == nil {
		if b == false {
			t.Log("Bool value for 'bool2' is false.")
		} else {
			t.Error("Bool value for 'bool2' is not false.")
		}
	} else {
		t.Error("Error getting bool value from property 'bool2':", err)
	}


	b, err = properties.Bool("level2.bool3")
	if err == nil {
		if b == true {
			t.Log("Bool value for 'bool3' is true.")
		} else {
			t.Error("Bool value for 'bool3' is not true.")
		}
	} else {
		t.Error("Error getting bool value from property 'bool3':", err)
	}


	b, err = properties.Bool("level2", "bool4")
	if err == nil {	
		if b == false {
			t.Log("Bool value for 'bool4' is false.")
		} else {
			t.Error("Bool value for 'bool4' is not false.")
		}
	} else {
		t.Error("Error getting bool value from property 'bool4':", err)
	}


	b, err = properties.Bool("level2.bool5[0]")
	if err == nil {
		if b == true {
			t.Log("Bool value for 'bool5[0]' is true.")
		} else {
			t.Error("Bool value for 'bool5[0]' is not true.")
		}
	} else {
		t.Error("Error getting bool value from property 'bool5[0]':", err)
	}
	

	b, err = properties.Bool("level2", "bool5", 1)
	if err == nil {
		if b == false {
			t.Log("Bool value for 'bool5[1]' is false.")
		} else {
			t.Error("Bool value for 'bool5[1]' is not false.")
		}
	} else {
		t.Error("Error retrieving bool value from property 'bool5[1]':", err)
	}
	
	b = properties.BoolDefault(true, "level2", "bool6")
	if b == true {
		t.Log("Default Bool value for 'bool6' is true.")
	} else {
		t.Error("Default Bool value for 'bool6' is not true.")
	}
}

