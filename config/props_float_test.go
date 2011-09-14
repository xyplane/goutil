//
//
//
package config_test

import(
	"config"
	"strings"
	"testing"
)

var TestFloatConfigData = `{
	"float1":1,
	"float2":2.0,
	"level2":{
		"float3":3.0e-3,
		"float4":[ 1.0, 2.0 ]
	}
}`

func TestFloat(t *testing.T) {

	t.Log("Read the following JSON config data:\n" + TestFloatConfigData)

	properties, err := config.ReadProperties(strings.NewReader(TestFloatConfigData))
	if err == nil {
		t.Log("Success reading config properties.")
	} else {
		t.Fatal("Error reading config properties:", err)
	}
	
	var f float64

	f, err = properties.Float64("float1")
	if err == nil {
		if f == 1.0 {
			t.Log("Float64 value for 'float1' is 1.0.")
		} else {
			t.Error("Float64 value for 'float1' is not 1.0.")
		}
	} else {
		t.Error("Error getting Bool value from property 'float1':", err)
	}

	f, err = properties.Float64("float2")
	if err == nil {
		if f == 2.0 {
			t.Log("Float64 value for 'float2' is 2.0.")
		} else {
			t.Error("Float64 value for 'float2' is not 2.0.")
		}
	} else {
		t.Error("Error getting Bool value from property 'float2':", err)
	}

	f, err = properties.Float64("level2", "float3")
	if err == nil {
		if f == 3.0e-3 {
			t.Log("Float64 value for 'float3' is 0.003.")
		} else {
			t.Error("Float64 value for 'float3' is not 0.003.")
		}
	} else {
		t.Error("Error getting Bool value from property 'float3':", err)
	}
	
	f, err = properties.Float64("level2.float4[0]")
	if err == nil {
		if f == 1.0 {
			t.Log("Float64 value for 'float4[0]' is 1.0.")
		} else {
			t.Error("Float64 value for 'float4[0]' is not 1.0.")
		}
	} else {
		t.Error("Error getting Bool value from property 'float4[0]':", err)
	}
	
	f, err = properties.Float64("level2", "float4", 1)
	if err == nil {
		if f == 2.0 {
			t.Log("Float64 value for 'float4[1]' is 2.0.")
		} else {
			t.Error("Float64 value for 'float4[1]' is not 2.0.")
		}
	} else {
		t.Error("Error getting Bool value from property 'float4[1]':", err)
	}
	
	f = properties.Float64Default(1.0, "level2", "float5")
	if err == nil {
		if f == 1.0 {
			t.Log("Default Float64 value for 'float5' is 1.0.")
		} else {
			t.Error("Default Float64 value for 'float5' is not 1.0.")
		}
	} else {
		t.Error("Error getting Bool value from property 'float5':", err)
	}
}
