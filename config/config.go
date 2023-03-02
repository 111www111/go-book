package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var YamlConfig Yaml

type Yaml struct {
	Server struct {
		Port string `yaml:"port"`
	}
	Mysql struct {
		Driver   string `yaml:"driver"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
	Redis struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
}

func Config() {
	//读取根目录config文件
	ymlFile, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(ymlFile, &YamlConfig)
	if err != nil {
		panic(err)
	}
}
