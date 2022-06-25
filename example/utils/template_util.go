package utils

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/BurntSushi/toml"
)

// Config 代表 TOML 配置文件的结构
type Config struct {
	Database DatabaseConfig
}

// DatabaseConfig 代表与数据库相关的配置选项
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func main() {
	// 加载 TOML 文件
	var config Config
	if _, err := toml.DecodeFile("/Users/menghuawei/IdeaProjects/my-project/wei-notebook/Mixinfo/template/example.toml", &config); err != nil {
		log.Fatalf("解码 TOML 文件时出错: %v", err)
	}

	// 加载 YAML 模板文件
	tmpl, err := template.ParseFiles("/Users/menghuawei/IdeaProjects/my-project/wei-notebook/Mixinfo/template/example.yml.tmpl")
	if err != nil {
		log.Fatalf("解析模板文件时出错: %v", err)
	}

	// 创建一个文件来保存渲染后的模板
	outputFile, err := os.Create("/Users/menghuawei/IdeaProjects/my-project/wei-notebook/Mixinfo/template/example.yml")
	if err != nil {
		log.Fatalf("创建输出文件时出错: %v", err)
	}
	defer outputFile.Close()

	// 使用解析后的配置执行模板
	if err := tmpl.Execute(outputFile, config); err != nil {
		log.Fatalf("执行模板时出错: %v", err)
	}

	fmt.Println("模板成功渲染到 example.yml")
}
