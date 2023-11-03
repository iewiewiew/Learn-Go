package main

/**
 * @author       weimenghua
 * @time         2023-11-24 10:07
 * @description  mapstructure 是一个 Go 语言库，用于将一个 map[string]interface{} 对象映射到一个结构体中。这个库非常有用，因为它可以将一个未知结构的 map 数据转换为一个已知结构的 Go 语言对象。
 */

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type AppConfig struct {
	Redis RedisConfig `mapstructure:"redis"`
}

func main() {
	// 读取配置文件
	file, err := os.Open("config.yaml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	// 将配置文件中的数据映射到结构体中
	var config AppConfig
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	// 输出配置信息
	fmt.Println("Redis Addr:", config.Redis.Addr)
	fmt.Println("Redis Username:", config.Redis.Username)
	fmt.Println("Redis Password:", config.Redis.Password)
	fmt.Println("Redis DB:", config.Redis.DB)
}
