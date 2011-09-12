package config

import (

)

type ConfigFile struct {
	fname string
	props Properties
}

func ReadConfigFile(fname string) (c *ConfigFile, err os.Error) {

	os.


}

func (c *ConfigFile) SetFileName(fname string) {
	c.fname = fname
}

func (c *ConfigFile) WriteConfigFile() os.Error {


} 

