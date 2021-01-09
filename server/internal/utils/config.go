package utils

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

//DBConfig is a struct for storing database information
type DBConfig struct {
	DBName     string `yaml:"dbname"`
	DBPassword string `yaml:"password"`
	DBUser     string `yaml:"user"`
}

type jwtSecret struct {
	secret string `yaml:"jwt_secret"`
}

//GetDBConfig parses database config information
func GetDBConfig(path string) DBConfig {
	var c DBConfig
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Can't load config file")
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatal("Can't load config file")
	}
	return c
}

//GetJWTSecret provides with jwt secret key
func GetJWTSecret() ([]byte, error) {
	path := "../../config/jwt_secret.yaml"
	var c jwtSecret
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}, err
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return []byte{}, err
	}
	return []byte(c.secret), nil
}
