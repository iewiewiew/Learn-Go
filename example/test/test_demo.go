package test

import "testing"

/**
# 默认使用goproxy.cn用户可手动调整
export GOPROXY=https://goproxy.cn
# 默认的单元测试命令
# 输出测试报告目录到当前工作目录,可自动上传并展示
mkdir -p golang-report
# 未使用Go Mod的用户需要打开一下注释 注：如果执行一些命令报错，则初始化项目：go mod init exmaple
sudo go test -short -v -json -cover -coverprofile ../golang-report/cover.out ./... > ../golang-report/report.jsonl
sudo go tool cover -html=../golang-report/cover.out -o ../golang-report/index.html
*/

func Add1(a, b int) int {
	return a + b
}

// TestAdd1 成功用例
func TestAdd1(t *testing.T) {
	r := Add1(1, 2)
	if r != 3 {
		t.Errorf("Add(1,2) failed. Got %d, expected 3.", r)
	}
}

// TestAdd2 跳过用例 go test -short
func TestAdd2(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	r := Add1(1, 2)
	if r != 2 {
		t.Errorf("Add(1,2) failed. Got %d, expected 3.", r)
	}
}

// TestAdd3 失败用例
func TestAdd3(t *testing.T) {
	r := Add1(1, 2)
	if r != 1 {
		t.Errorf("Add(1,2) failed. Got %d, expected 3.", r)
	}
}
