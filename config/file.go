package config

import (

)

type ConfigFile struct {
	Properties
	fname string
}

func (c *ConfigFile) SetFileName(fname string) {
	c.fname = fname
}
