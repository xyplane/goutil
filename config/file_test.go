// Copyright 2011 Dylan Maxwell.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package config_test

import (
	"config"
	"testing"
)

var TestFileName = "testdata/config.json"

func TestFile(t *testing.T) {

	t.Log("Read config properties from file:", TestFileName)

	c, err := config.ReadConfigFile(TestFileName)
	if err == nil {
		t.Log("Success reading test config file")
	} else {
		t.Fatal("Error reading test config file:", err)
	}

	var host string
	host, err = c.String("host")
	if err == nil {
		if host == "localhost" {
			t.Log("String value for property 'host' is 'localhost'.")
		} else {
			t.Error("String value for property 'host' is not 'localhost'.")
		}
	} else {
		t.Error("Error getting string value from property 'host':", err)
	}
}
