package main

import (
	"github.com/NameLessCorporation/user-grpc-server/internal/models"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"

	"github.com/NameLessCorporation/user-grpc-server/internal/app"
)

const configLevel = "CONFIG_LEVEL"
const configPath = "CONFIG_PATH"

const defaultConfigPath = "./configs"

const localConfigPath = "/local.yaml"
const prodConfigPath = "/prod.yaml"

func main() {
	configPath := getConfigPath()
	config, err := getConfig(configPath)
	if err != nil {
		log.Printf("package main: config error \n%v", err)
	}

	app.Run(config)
}

func getConfigPath() string {
	configPath := os.Getenv(configPath)
	if configPath == "" {
		configPath = defaultConfigPath
	}

	envConfigLevel := os.Getenv(configLevel)

	if envConfigLevel == "LOCAL" || envConfigLevel == "" {
		configPath += localConfigPath
	} else if envConfigLevel == "PROD" {
		configPath += prodConfigPath
	}

	return configPath
}

func getConfig(path string) (*models.Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config *models.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
