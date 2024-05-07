package conf

import "fmt"

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

// 如果配置映射成Config对象

// 从Toml格式的配置文件加载配置
func LoadConfigFromToml(filePath string) error {
	config = NewDefaultConfig()

	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("load config from toml file error: %v", err)
	}
	return nil
}

// 从环境变量加载配置
func LoadConfigFromEnv() error {
	config = NewDefaultConfig()
	return env.Parse(config)
}

// 加载全局配置
func loadConfig() (err error) {
	// 加载全局db
	db, err = config.MySQL.getDBConn()
	if err != nil {
		return fmt.Errorf("load global db error: %v", err)
	}
	return nil
}
