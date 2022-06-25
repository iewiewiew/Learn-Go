package chapter08_middle_software

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

/**
 * @author       weimenghua
 * @time         2024-01-15 11:24
 * @description  读取 Kafka
 */

func main() {
	// Kafka连接配置
	config := kafka.WriterConfig{
		Brokers: []string{"localhost:9092"}, // Kafka集群的地址
		Topic:   "my-topic",                 // Kafka主题名称
	}

	// 创建Kafka写入器
	writer := kafka.NewWriter(config)

	// 发送消息到Kafka主题
	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("key1"),
			Value: []byte("Hello, Kafka!"),
		},
		kafka.Message{
			Key:   []byte("key2"),
			Value: []byte("Welcome to Kafka!"),
		},
	)
	if err != nil {
		log.Fatal("无法发送消息到Kafka:", err)
	}

	fmt.Println("消息已成功发送到Kafka主题")

	// 关闭Kafka写入器
	writer.Close()

	// 创建Kafka读取器
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"}, // Kafka集群的地址
		Topic:   "my-topic",                 // Kafka主题名称
		GroupID: "my-consumer-group",        // 消费者组ID
	})

	// 从Kafka主题中读取消息
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("无法读取消息:", err)
		}

		fmt.Printf("接收到消息: Key = %s, Value = %s\n", string(msg.Key), string(msg.Value))
	}

	// 关闭Kafka读取器
	reader.Close()
}
