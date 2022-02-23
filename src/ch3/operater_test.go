package ch3_test

import "testing"

// go没有前置++/--，如++1/--1
func TestPreAuto(t *testing.T) {
	a := 1
	// t.Log(++a)  //报错：expected operand, found '++'
	a--
	a++
	t.Log(a)

}

func TestCmpareArray(t *testing.T) {

	// 数组比较需要维度及长度一样
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	// t.Log(a == c) //报错，长度不一致
	t.Log(a == d)
}

//位运算-按位清零&^
const (
	Readable   = 1 << iota //0001
	Writable               //0010
	Executable             //0100
)

func TestBitClear(t *testing.T) {
	a := 7            //0111
	a = a &^ Readable //按位清零&^运算
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
