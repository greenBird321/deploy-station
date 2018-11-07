package env

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"path/filepath"
	"log"
)

type MysqlConf struct {
	Host   string `yaml:"host"`
	Port   int16  `yaml:"port"`
	Dbname string `yaml:"dbname"`
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
}

type ENV struct {
	Mysql  MysqlConf `yaml:"mysql"`
	Domain string    `yaml:"domain"`
}

func AppEnv() (ENV, error) {
	t := ENV{}

	fileAbs, _ := filepath.Abs("./config/app.yml")
	content, err := ioutil.ReadFile(fileAbs)
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal(content, &t)
	if err != nil {
		log.Println(err.Error())
	}

	return t, nil
}
