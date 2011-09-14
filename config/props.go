// Copyright 2011 Dylan Maxwell.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package config

import (
	"os"
	"io"
	"fmt"
	"json"
	"regexp"
	"strings"
	"strconv"
	"reflect"
)

var PropNameDelim = "."
var PropNameRegex = regexp.MustCompile("^(.+)\\[([0-9]+)\\]$")

type Properties struct {
	root interface{}
}

// ReadProperties decodes JSON data and stores it in a Properties structure.
func ReadProperties(r io.Reader) (*Properties, os.Error) {
	var root interface{}
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&root)
	if err != nil {
		return nil, err
	}
	return &Properties{root}, nil
}

// Bool retrieves a boolean property value and an error if not found.
func (p *Properties) Bool(name ...interface{}) (bool, os.Error) {
	prop, err := p.Property(name...)
	if err != nil {
		return false, err
	}
	v, ok := prop.(bool)
	if !ok {
		err = os.NewError("property is not of type 'bool'.")
		return false, err
	}
	return v, nil
}

// BoolDefault retrieves a boolean property value or the specified default.
func (p *Properties) BoolDefault(dflt bool, name ...interface{}) bool {
	v, err := p.Bool(name...)
	if err != nil {
		return dflt
	}
	return v
}

// Int64 retrieves a int64 property value or an error if not found.
func (p *Properties) Int64(name ...interface{}) (int64, os.Error) {
	f, err := p.Float64(name...)
	if err != nil {
		return 0, err
	}
	return int64(f), nil
}

// Int64Default retrieves a int64 property value or the specified default.
func (p *Properties) Int64Default(dflt int64, name ...interface{}) int64 {
	v, err := p.Float64(name...)
	if err != nil {
		return dflt
	}
	return int64(v)
}

// Float64 retrieves a float64 property value or an error if not found.
func (p *Properties) Float64(name ...interface{}) (float64, os.Error) {
	prop, err := p.Property(name...)
	if err != nil {
		return 0.0, err
	}
	v, ok := prop.(float64)
	if !ok {
		err = os.NewError("property is not of type 'float64'.")
		return 0.0, err
	}
	return v, nil
}

// Float64Default retrieves a float64 property value or the specified default.
func (p *Properties) Float64Default(dflt float64, name ...interface{}) float64 {
	v, err := p.Float64(name...)
	if err != nil {
		return dflt
	}
	return v
}

// String retrieves a string property value or an error if not found.
func (p *Properties) String(name ...interface{}) (string, os.Error) {
	prop, err := p.Property(name...)
	if err != nil {
		return "", err
	}
	v, ok := prop.(string)
	if !ok {
		err = os.NewError("property is not of type 'string'.")
		return "", err
	}
	return v, nil
}

// StringDefault retrieves a string property value or the specified default.
func (p *Properties) StringDefault(dflt string, name ...interface{}) string {
	v, err := p.String(name...)
	if err != nil {
		return dflt
	}
	return v
}

// Properties retrieves a Properties value or an error if not found.
func (p *Properties) Properties(name ...interface{}) (*Properties, os.Error) {
	prop, err := p.Property(name...)
	if err != nil {
		return nil, err
	}
	return &Properties{prop}, nil
}

// Property retrieves a raw Property value and an error if not found. 
func (p *Properties) Property(name ...interface{}) (interface{}, os.Error) {
	sname, err := coerce(name...)
	if err != nil {
		return nil, err
	}

	if len(sname) == 1 {
		sname = split(sname[0])
	}

	var cur interface{} = p.root
	for _, sn := range sname {
		switch v := cur.(type) {
		case map[string]interface{}:
			var ok bool
			cur, ok = v[sn]
			if !ok {
				err := os.NewError(fmt.Sprint("map property does not contain key: ", sn))
				return nil, err
			}
		case []interface{}:
			idx, err := strconv.Atoi64(sn)
			if err != nil {
				return nil, err
			}
			if (idx < 0) || (idx >= int64(len(v))) {
				err = os.NewError(fmt.Sprint("array property does not contain index: ", idx))
				return nil, err
			}
			cur = v[idx]
		}
	}
	return cur, nil
}

func split(name string) []string {
	var sname []string
	names := strings.Split(name, PropNameDelim)
	for _, n := range names {
		if len(n) > 0 {
			match := PropNameRegex.FindStringSubmatch(n)
			if match == nil {
				sname = append(sname, n)
			} else {
				sname = append(sname, match[1], match[2])
			}
		}
	}
	return sname
}

func coerce(name ...interface{}) (sname []string, err os.Error) {
L:
	for _, n := range name {
		switch v := n.(type) {
		case string:
			sname = append(sname, v)
		case func() string:
			sname = append(sname, v())
		case fmt.Stringer:
			sname = append(sname, v.String())
		case int:
			sname = append(sname, strconv.Itoa(v))
		case int64:
			sname = append(sname, strconv.Itoa64(v))
		case float32:
			sname = append(sname, strconv.Itoa64(int64(v)))
		case float64:
			sname = append(sname, strconv.Itoa64(int64(v)))
		default:
			err = os.NewError(fmt.Sprint("name cannot be coerced from type: ", reflect.TypeOf(n)))
			break L
		}
	}
	return
}
