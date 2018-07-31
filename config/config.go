package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

var conf *Config
var once sync.Once

type Config struct {
	DB        DB        `json:"database"`
	SecretKey SecretKey `json:"secretKey"`
}

type DB struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

type SecretKey struct {
	Key string `json:"key"`
}

func FromFile(path string) *Config {
	once.Do(func() {
		bytes, err := ioutil.ReadFile(path)

		if err != nil {
			panic("error reading config.json " + path + ": " + err.Error())
		}

		if err := json.Unmarshal(bytes, &conf); err != nil {
			panic("error parsing config.json " + path + ": " + err.Error())
		}
		fmt.Println(conf)
	})

	return conf
}

func (d *DB) DbConnURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Name)
}
