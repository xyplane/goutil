//
//
//
package config

import (
	"os"
)

type ConfigFile struct {
	*Properties
	fname string
}

// ReadConfigFile reads the specified file and reads the config properties.
func ReadConfigFile(fname string) (c *ConfigFile, err os.Error) {
	var f *os.File
	f, err = os.Open(fname)
	if err != nil {
		return
	}
	
	var p *Properties
	p, err = ReadProperties(f)
	if err != nil {
		return
	}
	
	return &ConfigFile{ p, fname }, nil
}

func (c *ConfigFile) SetFileName(fname string) {
	c.fname = fname
}
