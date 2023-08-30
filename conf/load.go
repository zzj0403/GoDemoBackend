package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

// LoadConfigFromToml todo
func LoadConfigFromToml(filePath string) error {
	config = NewDefaultConfig()
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("load config from file error, path:%s, %s", filePath, err)
	}

	return nil
}

// LoadConfigFromEnv todo
func LoadConfigFromEnv() {

}
