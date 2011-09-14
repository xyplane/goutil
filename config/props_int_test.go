//
//
//
package config_test

import(
	"config"
	"strings"
	"testing"
)

var TestIntConfigData = `{
	"int1":1,
	"int2":2.0,
	"level2":{
		"int3":3.0e-3,
		"int4":[ 1.01, 2.99 ]
	}
}`

func TestInt(t *testing.T) {

	t.Log("Read the following JSON config data:\n" + TestIntConfigData)

	properties, err := config.ReadProperties(strings.NewReader(TestIntConfigData))
	if err == nil {
		t.Log("Success reading config properties.")
	} else {
		t.Fatal("Error reading config properties:", err)
	}
	
	var i int64

	i, err = properties.Int64("int1")
	if err == nil {
		if i == 1 {
			t.Log("Int64 value for 'int1' is 1.")
		} else {
			t.Error("Int64 value for 'int1' is not 1.")
		}
	} else {
		t.Error("Error getting int64 value from property 'int1':", err)
	}
	
	i, err = properties.Int64("int2")
	if err == nil {
		if i == 2 {
			t.Log("Int64 value for 'int2' is 2.")
		} else {
			t.Error("Int64 value for 'int2' is not 2.")
		}
	} else {
		t.Error("Error getting int64 value from property 'int2':", err)
	}
	
	i, err = properties.Int64("level2.int3")
	if err == nil {
		if i == 0 {
			t.Log("Int64 value for 'int3' is 0.")
		} else {
			t.Error("Int64 value for 'int3' is not 0.")
		}
	} else {
		t.Error("Error getting int64 value from property 'int1':", err)
	}
	
	i, err = properties.Int64("level2.int4[0]")
	if err == nil {
		if i == 1 {
			t.Log("Int64 value for 'int4[0]' is 1.")
		} else {
			t.Error("Int64 value for 'int4[0]' is not 1.")
		}
	} else {
		t.Error("Error getting int64 value from property 'int4[0]':", err)
	}
	
	i, err = properties.Int64("level2", "int4", 1)
	if err == nil {
		if i == 2 {
			t.Log("Int64 value for 'int4[1]' is 2.")
		} else {
			t.Error("Int64 value for 'int4[1]' is not 2.")
		}
	} else {
		t.Error("Error getting int64 value from property 'int4[1]':", err)
	}
	
	i = properties.Int64Default(99, "level2", "int5");
	if i == 99 {
		t.Log("Int64 value for 'int5' is 99.")
	} else {
		t.Error("Int64 value for 'int5' is not 99.")
	}
}
