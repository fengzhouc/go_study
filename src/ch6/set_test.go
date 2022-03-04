package ch6_test

import "testing"

func TestMapForSet(t *testing.T) {
	myset := map[int]bool{}
	myset[1] = true //1 赋值了
	n := 2
	if myset[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	myset[3] = true
	t.Log(len(myset))
	delete(myset, 1) //删除元素
	n = 1
	if myset[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}
