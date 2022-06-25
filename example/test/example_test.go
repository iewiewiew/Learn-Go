package test

import (
	"fmt"
	"os"
	"testing"
)

/**
# 默认的单元测试命令，输出测试报告目录到当前工作目录，可自动上传并展示
go test -short
mkdir -p golang-report
sudo go test -short -v -json -cover -coverprofile ../golang-report/cover.out ./... > ../golang-report/report.jsonl
sudo go tool cover -html=../golang-report/cover.out -o ../golang-report/index.html
*/

func Sum(a, b int) int {
	return a + b
}

func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}

// TestSum1 成功用例
func TestSum1(t *testing.T) {
	r := Sum(1, 2)
	if r != 3 {
		t.Errorf("Sum(1,2) failed. Got %d, expected 3.", r)
	}
}

// TestSum2 跳过用例
func TestSum2(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	r := Sum(1, 2)
	if r != 2 {
		t.Errorf("Sum(1,2) failed. Got %d, expected 3.", r)
	}
}

// TestSum3 失败用例
func TestSum3(t *testing.T) {
	r := Sum(1, 2)
	if r != 1 {
		t.Errorf("Sum(1,2) failed. Got %d, expected 3.", r)
	}
}

// 如果测试文件中包含函数 TestMain，那么生成的测试将调用 TestMain(m)，而不是直接运行测试。调用 m.Run() 触发所有测试用例的执行，并使用 os.Exit() 处理返回的状态码，如果不为0，说明有用例失败。因此可以在调用 m.Run() 前后做一些额外的准备(setup)和回收(teardown)工作。
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
