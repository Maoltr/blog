package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	DB DB `json:"database"`
}

type DB struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

func FromFile(path string) *Config {
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		panic("error reading config.json " + path + ": " + err.Error())
	}

	var conf Config
	if err := json.Unmarshal(bytes, &conf); err != nil {
		panic("error parsing config.json " + path + ": " + err.Error())
	}
	fmt.Println(conf)

	return &conf
}

func (d *DB) DbConnURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Name)
}
