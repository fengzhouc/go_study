package ch5_test

import "testing"

//切片，略复杂的东西
//切片内部结构
//指针，len元素个数，cap内部数组的容量
func TestSliceInit(t *testing.T) {
	var s []int              //声明跟数组很像，唯一区别就是没有指定长度
	t.Log(s, len(s), cap(s)) //[] 0 0,len跟cap都是0,默认初始化为0

	s = append(s, 1)         //append填充值
	t.Log(s, len(s), cap(s)) //[1] 1 1

	//声明并赋值
	e := []int{1, 2, 3, 4}
	t.Log(e, len(e), cap(e)) //[1 2 3 4] 4 4
	//以上基本都是数组蕾丝
	//下面就是区别了
	f := make([]int, 3, 5)        //参数说明: type,len,cap,其中len个元素会初始化为默认值
	t.Log(f, len(f), cap(f))      //[0 0 0] 3 5,因为len为3，则f的值默认是只有3个值
	f = append(f, 1)              //这里填充就不是全部填充了，而是从未初始化的元素开始填充
	t.Log(f[0], f[1], f[2], f[3]) //虽然定义的容量是5，但是到这里只初始化了4个元素，所以不能访问f[4],会报index out of
	t.Log(f, len(f), cap(f))      //[0 0 0 1] 4 5
	//到这里大概能理解切片len跟cap的意义了吧，len指定默认初始化的元素个数，cap指定整个切片的大小
}

func TestSliceGrowing(t *testing.T) {
	//切片可变长
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, 1) //还有这里append，为什么不是直接append(s,1),而需要重新复制给s,就是因为切片的可变长,由于长度改变,地址也会变,所以需要重新赋值，所以这里就会触发很对GC了
		// append(s, 1) //会报错
		t.Log(s, len(s), cap(s)) //填充是依次增涨的，即len，但是cap确是翻倍增涨,1,2,4,8,16,当空间不够时触发增涨，如例子:1,2,4,4,8,8,8,8,16,16
	}
}

//共享结构体
func TestSliceMemory(t *testing.T) {
	year := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := year[3:6]
	t.Log(s1, len(s1), cap(s1)) //[4 5 6] 3 7,切片的截取有点不同，截取的值是一样的，但是cap却是3:len(year),就是到最后了，cap相当于[3:]
	s2 := year[5:9]
	t.Log(s2, len(s2), cap(s2)) //[6 7 8 9] 4 5
	s2[0] = 123                 //修改值
	t.Log(s1, len(s1), cap(s1)) //[4 5 123] 3 7
	t.Log(s2, len(s2), cap(s2)) //[123 7 8 9] 4 5
	//那所有这个地址的值都会改了，因为他们是共享year的,每个元素的值都是指向同一个地址的，一改全改
	t.Log(year, len(year), cap(year)) //[1 2 3 4 5 123 7 8 9 10] 10 10
}

func TestSliceCompare(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	if s1 == s2 { //报错了,提示切片只能与nil,即只能切片判空，而不能进行两个切片的比较
		t.Log(111)
	}
}
