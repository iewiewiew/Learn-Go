package test

import (
	"testing"
)

// go test -v
// 成功
func TestAdd(t *testing.T) {
	a := 10
	b := 20
	want := 30
	actual := Add(a, b)
	if want != actual {
		t.Errorf("Add函数参数:%d %d, 期望: %d, 实际: %d", a, b, want, actual)
	}
}

// 失败
func TestMul(t *testing.T) {
	a := 10
	b := 20
	want := 300
	actual := Mul(a, b)
	if want != actual {
		t.Errorf("Mul函数参数:%d %d, 期望: %d, 实际: %d", a, b, want, actual)
	}
}

// 成功
func TestDiv(t *testing.T) {
	a := 10
	b := 20
	want := 2
	actual := Div(a, b)
	if want != actual {
		t.Errorf("Div函数参数:%d %d, 期望: %d, 实际: %d", a, b, want, actual)
	}
}

// 跳过
func TestSkip(t *testing.T) {
	a := 10
	b := 20
	want := 30
	got := Add(a, b)
	if want == got {
		t.Skip("Skipping testing")
	}
}

// go test -bench=. @todo: 找不到测试文件
func BenchmarkCompute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}
