package ch4_test

import "testing"

//条件与循环

//循环
func TestLoop(t *testing.T) {

	//普通for循环，相比java少了（）
	// for i := 0; i < 10; i++ {
	// 	t.Log(i)
	// }

	//简化，相当于while (n<5)
	n := 0
	for n < 5 {
		n++
		t.Log(n)
	}

	//无限循环,相当于while (true)
	// n := 0
	// for {
	// 	n++
	// 	t.Log(n)
	// }
}

func TestCondition(t *testing.T) {
	//存在区别的点,除了常规的if condition，go还可以变量赋值
	if a := 3; a < 2 {
		t.Log(111)
	} else {
		t.Log(222)
	}
}

func TestSwitch(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i { //case默认是break，所以不需要显式break
		case 0, 2: //支持多个条件,只要一个命中即可
			t.Log("0,2", i)
		case 1, 3:
			t.Log("1,3", i)
		default:
			t.Log(i)
		}
	}
}

func TestCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch { //case如果写条件语句的话，那switch这里的i就不能写了，写了回报错
		// switch i { //这样写回报错
		case i%2 == 0: //case这里可以写条件语句
			t.Log("i%2==0", i)
		case i%2 == 1:
			t.Log("i%2==1", i)
		default:
			t.Log(i)
		}
	}
}
