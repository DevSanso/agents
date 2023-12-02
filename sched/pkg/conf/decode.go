package conf

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

func ReadTomlConfig(path string) (*Configure, error) {
	file,err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	
	var ret = &Configure{}
	_,err = toml.Decode(string(file), ret)
	
	if err != nil {
		ret = nil
		return nil, err
	}

	return ret,nil
}