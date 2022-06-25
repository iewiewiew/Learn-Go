package main

/**
 * @author       weimenghua
 * @time         2023-11-24 10:07
 * @description  mapstructure 是一个 Go 语言库,用于将一个 map[string]interface{} 对象映射到一个结构体中.这个库非常有用,因为它可以将一个未知结构的 map 数据转换为一个已知结构的 Go 语言对象.
 * Mapstructure是一个Go语言的包，用于将map类型的结构体或结构体类型的map进行转换。它可以将一个map转换为一个结构体，或者将一个结构体转换为一个map。使用mapstructure可以方便地将来自外部的数据源（如配置文件、数据库等）中的数据转换为结构化的Go数据类型，以便在程序中进行处理和使用。它通过定义一个结构体，并使用mapstructure标签指定map中的键与结构体字段的映射关系，从而实现map到结构体的转换。同样，也可以使用结构体转换为map的方式。
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
