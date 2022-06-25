package main

import "fmt"

/**
 * @author       weimenghua
 * @time         2024-07-19 15:55
 * @description
 */

type Server struct{}

// 为 Server 定义方法 aaa，返回 *Server 以便于方法链调用
func (s *Server) aaa() *Server {
	fmt.Println("aaa called")
	return s
}

// 为 Server 定义方法 bbb，返回 *Server 以便于方法链调用
func (s *Server) bbb() *Server {
	fmt.Println("bbb called")
	return s
}

func main() {
	// 创建一个 Server 实例
	server := &Server{}

	// 方法链调用
	// 在一个语句中连续调用多个返回相同类型的接收者方法。调用了 server.aaa().bbb()。由于 aaa 方法返回 *Server，我们可以直接在其后调用 bbb 方法。
	server.aaa().bbb()
}
