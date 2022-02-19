package ch1_test

import "testing"

func TestAutotype(t *testing.T) {
	// 普通变量定义
	// var a int = 1
	// var b int = 2

	// 集中定义
	// var (
	// 	a = 1
	// 	b = 2
	// )

	// :定义不需要类型，这里应该是go自己根据值来判断类型
	// a := 1
	// b := 2

	// 特性：自动类型推断,不需要类型，这里其实跟上面的:一样
	var a int = 1
	var b = 2

	t.Log(a, b)
}
