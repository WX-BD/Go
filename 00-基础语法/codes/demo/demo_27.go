package main

import (
	"fmt"
	"math"
)

/*
	type Member struct {
		id     int
		name   string
		email  string
		gender int
		age    int
	}
*/
var m2 = Member{1, "小明", "xiaoming@163.com", 1, 18} // 简短变量声明方式：m2 := Member{1,"小明","xiaoming@163.com",1,18}
var m3 = Member{id: 2, name: "小红"}                  // 简短变量声明方式：m3 := Member{id:2,"name":"小红"}

// 遍历字符串
func traversalString() {
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

func changeString() {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))
}

func sqrtDemo() {
	var a, b = 3, 4
	//var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c := int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func Change(m1 Member, m2 *Member) {
	m1.id = 1
	m1.name = "小明"
	m2.id = 2
	m2.name = "小绿"
	m2.email = "xiaolv@163.com"
	m2.age = 12
	m2.gender = 1
}

func main_() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	d1 := slice[6:8]
	fmt.Println(d1, len(d1), cap(d1))
	d2 := slice[:6:8]
	fmt.Println(d2, len(d2), cap(d2))
	/*
		d := [5]struct {
			x int
		}{}
		s := d[:]
		d[1].x = 10
		s[2].x = 20
		fmt.Println(d)
		fmt.Printf("%p, %p\n", &d, &d[0])
	*/
	/*
		m1 := Member{}
		m2 := new(Member)
		Change(m1, m2)
		fmt.Println(m1, m2)
	*/
	/*
		fmt.Println(m2.name)
		m3.name = "小花"
		fmt.Println(m3.name)
		age := &m3.age
		*age = 20
		fmt.Println(m3.age)
	*/
	//traversalString()
	//changeString()
	//sqrtDemo()
	//a := [2]int{1, 2}
	//println(len(a), cap(a))
	/*
		var arr1 [5]int
		var arr2 *[5]int
		printArr(&arr1)
		fmt.Println(&arr1)
		fmt.Printf("addr:%p\n", &arr1)
		fmt.Println(&arr2)
		fmt.Println(arr1)
	*/
	/*
		var arrPtr *[4]int           // 创建一个指针 arrPtr，指向一个数组
		var arr = [4]int{1, 2, 3, 4} // 创建一个数组并初始化
		arrPtr = &arr                // 将数组 arr的地址赋值给arrPtr
		fmt.Println("将 arr 的内存地址赋值给数组指针 arrPtr,   arrPtr=", arrPtr)
		fmt.Printf("arr 数组的地址为：%p\n", &arr)
		fmt.Printf("arrPtr 存储的地址为：%p\n", arrPtr)
		fmt.Println("(*arrPtr)[0] 通过指针访问数组的第一个元素：", (*arrPtr)[0])
		fmt.Println("arrPtr[0] 通过指针访问数组的第一个元素：", arrPtr[0])
	*/
	/*
		var ptrArr [4]*int
		a, b, c, d := 1, 2, 3, 4
		arr2 := [4]int{a, b, c, d} // 拷贝四个变量的值为函数组元素
		fmt.Println("数组 arr2 :", arr2)
		ptrArr = [4]*int{&a, &b, &c, &d} // 存的都是内存地址
		fmt.Println("指针数组 ptrArr :", ptrArr)
		fmt.Println(&a)
	*/
}
