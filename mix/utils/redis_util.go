package utils

/**
 * @author       weimenghua
 * @time         2023-08-15 14:36
 * @description  Go Redis 增删改查
 */

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

func go_redis() {
	// 创建 redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址和端口
		Password: "",               // Redis 服务器密码,如果没有密码则为空字符串
		DB:       0,                // Redis 数据库索引,默认为 0
	})

	// 创建连接
	// context.Background(): 这是一个上下文对象,通过调用 context.Background() 创建.上下文对象用于控制请求的行为,例如设置超时时间、取消请求等.在这个例子中,我们使用了默认的上下文背景.
	// Result(): 这是执行 Redis 命令并等待结果的方法.在这个例子中,Result() 方法将等待 Ping 命令的执行结果,并返回响应内容.
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Redis 连接成功!", pong)

	// 查询所有的 Key
	keys, err := client.Keys(context.Background(), "*").Result()

	if err != nil {
		log.Fatal(err)
	}

	// 打印所有的 Key
	fmt.Println("所有的 Key: ")
	for _, key := range keys {
		fmt.Println(key)
	}

	// 设置键值对
	err = client.Set(context.Background(), "mykey", "myvalue", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	// 查询单个 key
	value, err := client.Get(context.Background(), "mykey").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("key 的值：", value)

	// 删除单个 key
	result, err := client.Del(context.Background(), "mykey").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("删除 key 成功", result)

	// 关闭连接
	err = client.Close()
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	go_redis()
}
