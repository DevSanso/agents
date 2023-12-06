package conf

import (
	"github.com/BurntSushi/toml"
	"os"
)

func ReadTomlConfig(path string) (*Configure, error) {
	file,err := os.ReadFile(path)
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