package ch5_test

import "testing"

func TestArrayOne(t *testing.T) {
	// 可以看到go的数组定义其实跟Java很类似的，只是顺序顺序有点变，[]表示在类型前面了，Java是在后面
	var a [3]int //声明并初始化默认0值
	t.Log(a)     //[0 0 0]
	a[0] = 1     //重赋值
	t.Log(a)     //[1 0 0]

	b := [3]int{1, 2, 3} //声明并同时自定义初始化，:=这种方式不能默认初始化哈，会报错
	t.Log(b)

	c := [2][2]int{{1, 2}, {3, 4}} //二维数组
	t.Log(c)
	var d [2][2]int
	t.Log(d)

	e := [...]int{1, 2, 3, 4} //可以不具体指定长度，让其根据具体初始化的长度自行设置
	t.Log(e, len(e))          //[1 2 3 4] 4
}

func TestArrayTrave(t *testing.T) {
	//数组遍历,for
	a := [...]int{1, 2, 3, 4}
	for i := 0; i < len(a); i++ {
		t.Log(a[i])
	}

	//跟python很像的一个写法
	for index, v := range a {
		t.Log(index, v)
	}
	//但是呢，上面会有两个值，如果不全部使用，go编译会报错，那怎么办呢
	for _, v := range a { //_相当于占位，意味着我们并不关心这个值，所以就忽略了,当然这样也就不能使用这个值了
		t.Log(v)
	}
}

func TestArraySub(t *testing.T) {
	//数组截取
	a := [...]int{1, 2, 3, 4}
	t.Log(a[1:2]) //[2],不包含后边界的值
	t.Log(a[1:])  //[2 3 4],从1到最后
	t.Log(a[:3])  //[1,2,3],从0到3
	// t.Log(a[1::2])//跟python相似的，不过只有始末位置的制定，并没有步长的制定，这里会报错
	// t.Log(a[:-1])//还有想python可以写负值，然后会的倒序的数组，但go不支持，会报错
}
