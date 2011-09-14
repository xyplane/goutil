// Copyright 2011 Dylan Maxwell.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package config_test

import (
	"config"
	"strings"
	"testing"
)

var TestPropsConfigData = `{
	"string1":"Hello World",
	"level2":{
		"float2":3.0e-3,
		"mixed3":[ "one", 10 ]
	}
}`

func TestProps(t *testing.T) {

	t.Log("Read the following JSON config data:\n" + TestPropsConfigData)

	properties, err := config.ReadProperties(strings.NewReader(TestPropsConfigData))
	if err == nil {
		t.Log("Success reading config properties.")
	} else {
		t.Fatal("Error reading config properties:", err)
	}

	var p *config.Properties

	p, err = properties.Properties("string1")
	if err == nil {
		var s string
		s, err = p.String()
		if err == nil {
			if s == "Hello World" {
				t.Log("String value for Properties('string1').String() is 'Hello World'.")
			} else {
				t.Error("String value for Properties('string1').String() is not 'Hello World'.")
			}
		} else {
			t.Error("Error getting string value from property '':", err)
		}
	} else {
		t.Error("Error getting Properties value from property 'string1':", err)
	}

	p, err = properties.Properties("level2")
	if err == nil {
		var f float64
		f, err = p.Float64("float2")
		if err == nil {
			if f == 3.0e-3 {
				t.Log("Float64 value for Properties('level2').Float64('float2') is 3.0e-3.")
			} else {
				t.Error("Float64 value for Properties('level2').Float64('float2') is not 3.0e-3.")
			}
		} else {
			t.Error("Error getting float64 value from property 'float2':", err)
		}
	} else {
		t.Error("Error getting Properties value from property 'level2':", err)
	}

	p, err = properties.Properties("level2.mixed3")
	if err == nil {
		var s string
		s, err = p.String(0)
		if err == nil {
			if s == "one" {
				t.Log("String value for Properties('level2.mixed3').String(0) is 'one'.")
			} else {
				t.Error("Float64 value for Properties('level2.mixed3').String(0) is not 'one'.")
			}
		} else {
			t.Error("Error getting string value from property '0':", err)
		}
	} else {
		t.Error("Error getting Properties value from property 'level2.mixed3':", err)
	}

	p, err = properties.Properties("level2", "mixed3", 1)
	if err == nil {
		var i int64
		i, err = p.Int64()
		if err == nil {
			if i == 10 {
				t.Log("Int64 value for Properties('level2.mixed3[1]').Int64() is 10.")
			} else {
				t.Error("Int64 value for Properties('level2.mixed3[1]').Int64() is not 10.")
			}
		} else {
			t.Error("Error getting int64 value from property '':", err)
		}
	} else {
		t.Error("Error getting Properties value from property 'level2.mixed3[1]':", err)
	}
}
