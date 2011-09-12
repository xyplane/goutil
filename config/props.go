//
//
//
package config

import (
	"regexp"
	"strings"
	"strconv"
)

var PropNameRegex = regexp.MustCompile("(.+)(\[[0-9]+\])?\.*") 

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
	err := json.NewDecoder(r).Decode(root)
	if err != nil {
		return
	}
	p = Properties{ root }
	return
}


func (p Properties) Bool(name ...interface{}) (v bool, err os.Error) {
	prop, ok := p.Property(name)
	if !ok {
		return false, false;
	}
	
	return prop.(bool)
}

func (p Properties) BoolDefault(bool dflt, name ...interface{}) {
	

} 

func (p Properties) Property(name ...interface{}) (interface{}, bool) {

	var sname []string
	
	for(n := range name) {
		sn, coerceErr := coerceName(n)
		if coerceErr != nil {
			return 
		}
		append(sname, sn)
	}
	
	if len(sname) == 1 {
		sname = splitName(sname)
	}
	
	if len(sname) == 0 {
	
	}
		
	var current interface{} = p
	
	for sn := range sname {
		switch i := current.(type) {
		case map[string]interface{}:
			var ok bool
			current, ok = i[sn] 
			if !ok {
				return
			}
			
		case []interface{}
			idx, err = strconv.Atoi64(sn)
			if err != nil {
				return
			}
			if idx >= len(i) {
				return
			}
			current = i[idx]
		}
	}
	
	return current
}

func splitName(name string) []string {
	var names []string
	matches := NameRegex.FindAllSubmatch(name)	
	for(match := range matches) {
		append(names, match[1])
		if match[2] != "" {
			append(names, match[2])
		}
	}
	return names;
} 

func coerceName(name interface{}) (sname string, err os.Error) {
	switch i := name.(type) {
		case int64:
			sname = strconv(i)
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
