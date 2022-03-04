package ch6_test

import "testing"

func TestMapInit(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	m1 := map[string]int{}
	m1["one"] = 1
	m2 := make(map[string]int, 10 /*init cap*/) //没有len，因为map是键值对，没办法初始化为0
	t.Log("m:", m)
	t.Log("m1:", m1)
	t.Log("m2:", m2)
}

//key不存在时，go会初始化为0
func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	// m1[3] = 0  //不赋值，则下面的条件就是得到key 3不存在

	if v, ok := m1[3]; ok { //条件新写法？？v, ok := m1[3]，v为值，ok为状态，是否成功
		t.Logf("key 3's value is %d", v)
	} else {
		t.Log("key 3 is not existsing", v) //所以不能通nil来判断是否存在某元素
	}
}

//遍历
func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
	for k, v := range m1 {
		t.Log(k, v)
	}
}

//map写工程模式，其实就是value不一定是值，也可以是函数
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }      //op的平方
	m[3] = func(op int) int { return op * op * op } //op的立方
	t.Log(m[1](2), m[2](2), m[3](2))                //取map值调用函数

}
