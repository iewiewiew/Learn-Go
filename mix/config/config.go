package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

/**
 * @author       weimenghua
 * @time         2024-07-12 10:30
 * @description
 */

var Conf Config

type Config struct {
	Database struct {
		UserName string `yaml:"user_name"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		DBName   string `yaml:"dbname"`
		Port     string `yaml:"port"`
	} `yaml:"database"`
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
}

func init() {
	configPath := "config/config.yaml"
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("Failed to unmarshal config file: %v", err)
	}
	log.Println("Config loaded successfully")
}

func LoadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	configPath := "config.yaml"
	config, err := LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Printf("Database UserName:%s\n", config.Database.DBName)
}
