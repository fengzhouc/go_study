package ch2_test

import "testing"

/*
基本类型
bool
string
int,int8,int16,int32,int64
uint,uint8,uint16,uint32,uint64,uintptr
byte  // uint8的别名，其实就是uint8类型
rune // uint32的别名，其实就是uint32类型
float32,float64
complex64,complex128
*/

// go特点：不支持隐式类型转换

func TestDatatype(t *testing.T) {
	//type myint int64  // int64的别名myint，这样实际应该都是int64，但也不行，还是报错
	// var a myint = 1
	var a int = 1
	var b int64
	// b = a // 编译报错：cannot use a (type int) as type int64 in assignment
	b = int64(a) // 显式类型转换后就可以
	t.Log(a, b)
}

// go虽然支持指针，但不支持指针运算
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a //获取指针的方式跟c/c++一样，&取址符
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
	// 指针运算会报错
	// aPtr = aPtr + 1 //报错：invalid operation: aPtr + 1 (mismatched types *int and int)
}

func TestString(t *testing.T) {
	var s string         //
	t.Log("*" + s + "*") //输出:**,为空值
}
