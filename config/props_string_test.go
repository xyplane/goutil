// Copyright 2011 Dylan Maxwell.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package config_test

import (
	"config"
	"strings"
	"testing"
)

var TestStringConfigData = `{
	"string1":"",
	"string2":"HW!",
	"level2":{
		"string3":"3.0e-3",
		"string4":[ "one", "two" ]
	}
}`

func TestString(t *testing.T) {

	t.Log("Read the following JSON config data:\n" + TestStringConfigData)

	properties, err := config.ReadProperties(strings.NewReader(TestStringConfigData))
	if err == nil {
		t.Log("Success reading config properties.")
	} else {
		t.Fatal("Error reading config properties:", err)
	}

	var s string

	s, err = properties.String("string1")
	if err == nil {
		if s == "" {
			t.Log("String value for 'string1' is ''.")
		} else {
			t.Error("String value for 'string1' is not ''.")
		}
	} else {
		t.Error("Error getting string value from property 'string1':", err)
	}

	s, err = properties.String("string2")
	if err == nil {
		if s == "HW!" {
			t.Log("String value for 'string2' is 'HW!'.")
		} else {
			t.Error("String value for 'string2' is not 'HW!'.")
		}
	} else {
		t.Error("Error getting string value from property 'string2':", err)
	}

	s, err = properties.String("level2.string3")
	if err == nil {
		if s == "3.0e-3" {
			t.Log("String value for 'string3' is '3.0e-3'.")
		} else {
			t.Error("String value for 'string3' is not '3.0e-3'.")
		}
	} else {
		t.Error("Error getting string value from property 'string3':", err)
	}

	s, err = properties.String("level2.string4[0]")
	if err == nil {
		if s == "one" {
			t.Log("String value for 'string4[0]' is 'one'.")
		} else {
			t.Error("String value for 'string4[0]' is not 'one'.")
		}
	} else {
		t.Error("Error getting string value from property 'string4[0]':", err)
	}

	s, err = properties.String("level2.string4[1]")
	if err == nil {
		if s == "two" {
			t.Log("String value for 'string4[1]' is 'two'.")
		} else {
			t.Error("String value for 'string4[1]' is not 'two'.")
		}
	} else {
		t.Error("Error getting string value from property 'string4[1]':", err)
	}

	s = properties.StringDefault("test", "level2", "string5")
	if s == "test" {
		t.Log("Default String value for 'string5' is 'test'.")
	} else {
		t.Error("Default String value for 'string5' is not 'test'.")
	}
}
