package ch7_test

import "testing"

/*
string是数据类型，前面也提到了，默认初始化为0，不是引用或指针类型
string是只读的byte slice，len函数可以它所包含的byte数
string的byte数组可以存放任何数据
*/

func TestString(t *testing.T) {
	var s string
	t.Log(s, len(s)) //默认初始化为0
	s = "Hello"
	t.Log(s, len(s)) //计算的byte长度
	// s[2] = "1" //string是只读切片，不能重新赋值
	t.Log(s[2])
	s = "\xe4\xb8\xa5" //汉字严的十六进制
	// s = "\xe4\xba\xb5\xff" //最后\xff也会打印
	t.Log(s) //这里会自动转码为汉字
	s = "中"
	t.Log(s, len(s)) //len计算的byte数

	c := []rune(s)
	// t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}
