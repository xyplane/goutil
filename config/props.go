//
//
//
package config

import (
	"os"
	"io"
	"fmt"
	"json"
	"regexp"
	"strconv"
)

var PropNameRegex = regexp.MustCompile("(.+)(\\[[0-9]+\\])?\\.*") 

type Properties struct {
	root interface{}
	//Bool(name ...interface{}) (bool, bool)
	//Int64(name string) (int64, bool)
	//String(name string) (string, bool)
	//Float64(name string) (float64, bool)
	//Property(name string) (interface{}, bool)
	//Properties(name string) (Properties, bool)
}


// ReadProperties decodes JSON data and stores it in a Properties structure.
func ReadProperties(r io.Reader) (p Properties, err os.Error) {
	var root interface{}	
	err = json.NewDecoder(r).Decode(root)
	if err != nil {
		return
	}
	p = Properties{ root }
	return
}

// Bool retrieves a boolean property value and an error if not found.
func (p Properties) Bool(name ...interface{}) (bool, os.Error) {
	i, err := p.Property(name)
	if err != nil {
		return false, err
	}
	v, ok := i.(bool)
	if !ok {
		err = os.NewError("")
		return false, err
	} 
	return v, nil
}

// BoolDefault retrieves a boolean property value or the specified default.
func (p Properties) BoolDefault(dflt bool, name ...interface{}) bool {
	v, err := p.Bool(name)
	if err != nil {
		return dflt
	}
	return v
}

// Property retrieves a raw Property value and an error if not found. 
func (p Properties) Property(name ...interface{}) (interface{}, os.Error) {
	var sname []string
	for _, n := range name {
		sn, err := coerce(n)
		if err != nil {
			return nil, err 
		}
		sname = append(sname, sn)
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
				err := os.NewError("")
				return nil, err
			}
		case []interface{}:
			idx, err := strconv.Atoi64(sn)
			if err != nil {
				return nil, err
			}
			if idx >= int64(len(v)) {
				err = os.NewError("")
				return nil, err
			}
			cur = v[idx]
		}
	}
	return cur, nil
}

func split(name string) []string {
	var names []string
	matches := PropNameRegex.FindAllStringSubmatch(name, -1)	
	for _, match := range matches {
		names = append(names, match[1])
		if match[2] != "" {
			names = append(names, match[2])
		}
	}
	return names;
} 

func coerce(name interface{}) (sname string, err os.Error) {
	switch i := name.(type) {
		case int64:
			sname = strconv.Itoa64(i)
		case string:
			sname = i	
		case fmt.Stringer:
			sname = i.String()
		case func() string:
			sname = i()
		default:
			err = os.NewError("")
	}
	return
}
