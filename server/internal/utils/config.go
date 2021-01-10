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

var dbcfgPath = "../../config/dbconfig.yaml"
var dbcfgBuildPath = "./config/dbconfig.yaml"

var jwtSecretPath = "../../config/jwt_secret.yaml"
var jwtSecretBuildPath = "./config/jwt_secret.yaml"

//GetDBConfig parses database config information
func GetDBConfig() DBConfig {
	var c DBConfig
	var yamlFile []byte
	yamlFile, err := ioutil.ReadFile(dbcfgPath)
	if err != nil {
		yamlFile, err = ioutil.ReadFile(dbcfgBuildPath)
		if err != nil {
			log.Fatal("Can't load config file")
		}
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatal("Can't load config file")
	}
	return c
}

//GetDBTestConfig parses a config file for testing purposes
//there is a better way to do it than this
func GetDBTestConfig() DBConfig {
	var c DBConfig
	var yamlFile []byte
	yamlFile, err := ioutil.ReadFile("../../../config/dbtestconfig.yaml")
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
	var c jwtSecret
	var yamlFile []byte
	yamlFile, err := ioutil.ReadFile(jwtSecretPath)
	if err != nil {
		yamlFile, err = ioutil.ReadFile(jwtSecretBuildPath)
		if err != nil {
			return []byte{}, err
		}
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return []byte{}, err
	}
	return []byte(c.secret), nil
}
