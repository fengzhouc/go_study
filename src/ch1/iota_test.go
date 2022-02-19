package ch1_test

import "testing"

func TestIota(t *testing.T) {
	// iota关键字可以快速设置连续值
	// const (
	// 	a = 1 + iota
	// 	b
	// 	c
	// 	d
	// )
	// bit也可以使用，即位运算
	const (
		a = 1 << iota
		b
		c
		d
	)
	// 除了iota这个连续常量的特性，其他都跟Java一样

	t.Log(a, b, c, d)
}
