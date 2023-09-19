package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type student struct {
	id   int
	name string
	age  int
}

// Student 学生
type Student_ struct {
	ID     int `json:"id"`
	gender string
	Name   string `json:"name"`
}

// Class 班级
type Class struct {
	Title    string
	Students []*Student_
}

func main() {
	/*
		p9 := newPerson("pprof.cn", "测试", 90)
		fmt.Printf("%#v\n", p9)
	*/
	//testJSON()
	//mapTest3()
	//mapTest1()
	//mapTemp()
	//sliceMap()
	//mapGo()
	//scoreMap()
	//mapTest2()
	//arrTest1()
	//main_()
	//deleteMap()
	//sortMap()
	//testChan()
	//testGoroutine()
	//testChan1()
	//testSelect()
	//testChan2()
	//testFunc()
	//testDefer()
	//sortByLength()
	//testClosure2()
	testdefer2()
}

func testdefer2() {
	var whatever [5]struct{}
	for i := range whatever {
		defer func() { fmt.Println(i) }()
	}
}

/*
	func testFibonacci() {
		fibonaci := func(i int) int {
			if i == 0 {
				return 0
			}
			if i == 1 {
				return 1
			}
			return func(i-1) + func(i-2)
		}
	}
*/
func testDefer1() {
	var whatever [5]struct{}
	for i := range whatever {
		defer fmt.Println("defer:", i)
		fmt.Println("main:", i)
	}
}

func testClosure2() {
	test01 := func(base int) (func(int) int, func(int) int) {
		// 定义2个函数，并返回
		// 相加
		add := func(i int) int {
			base += i
			return base
		}
		// 相减
		sub := func(i int) int {
			base -= i
			return base
		}
		// 返回
		return add, sub
	}
	f1, f2 := test01(10)
	fmt.Println(f1(0), f2(0))
	// base一直是没有消
	fmt.Println(f1(1), f2(2))
	// 此时base是9
	fmt.Println(f1(3), f2(4))
}

func testClosure1() {
	var add = func(base int) func(int) int {
		return func(i int) int {
			base += i
			return base
		}
	}
	tmp1 := add(10)
	fmt.Println(tmp1(0))
	fmt.Println(tmp1(1), tmp1(2))

	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2))
}

func testClosure() func() {
	x := 100
	fmt.Printf("x1 (%p) = %d\n", &x, x)

	return func() {
		fmt.Printf("x2 (%p) = %d\n", &x, x)
	}
}

func sortByLength() {
	strings := []string{"apple", "banana", "cherry", "date"}

	sort.Slice(strings, func(i, j int) bool {
		return len(strings[i]) > len(strings[j])
	})

	fmt.Println(strings)

}

func testFunc1() {
	// --- function variable ---
	fn := func() { println("Hello, World!") }
	fn()

	// --- function collection ---
	fns := [](func(x int) int){
		func(x int) int { return x + 1 },
		func(x int) int { return x + 2 },
	}
	println(fns[0](100))

	// --- function as field ---
	d := struct {
		fn func() string
	}{
		fn: func() string { return "Hello, World!" },
	}
	println(d.fn())

	// --- channel of function ---
	fc := make(chan func() string, 2)
	fc <- func() string { return "Hello, World!" }
	println((<-fc)())
}

func testDefer() {
	var add = func(x, y int) (z int) {
		defer func() {
			println("defer:", z)
		}()

		z = x + y
		return z + 200
	}
	println(add(1, 2))
}

func testFunc() {
	var test = func(s string, n ...int) string {
		var x int
		for _, i := range n {
			x += i
		}

		return fmt.Sprintf(s, x)
	}
	s := []int{1, 2, 3}
	res := test("sum: %d", s...) // slice... 展开slice
	println(res)
}

func testChan2() {
	var counter = func(out chan<- int) {
		for i := 0; i < 100; i++ {
			out <- i
		}
		close(out)
	}
	var squarer = func(out chan<- int, in <-chan int) {
		for i := range in {
			out <- i * i
		}
		close(out)
	}
	var printer = func(in <-chan int) {
		for i := range in {
			fmt.Println(i)
		}
	}
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}

func testSelect() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}

func testChan1() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}

func testGoroutine() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine: ", i)
			time.Sleep(time.Second)
		}
	}()
	// 在主 Goroutine 中执行这个循环
	for i := 0; i < 5; i++ {
		fmt.Println("Main: ", i)
		time.Sleep(time.Second)
	}
}

func testChan() {
	ch := make(chan int, 1)
	ch <- 10
	fmt.Println("发送成功！", ch)
	close(ch)
}

func sortMap() {
	map1 := make(map[int]string, 5)
	map1[1] = "www.topgoer.com"
	map1[2] = "rpc.topgoer.com"
	map1[5] = "ceshi"
	map1[3] = "xiaohong"
	map1[4] = "xiaohuang"
	sli := []int{}
	for k, _ := range map1 {
		sli = append(sli, k)

	}
	fmt.Println(sli)
	sort.Ints(sli)
	fmt.Println(sli)
	for i := 0; i < len(map1); i++ {
		fmt.Println(map1[sli[i]])
	}
}

func deleteMap() {
	ce := make(map[int]student)
	ce[1] = student{1, "xiaolizi", 22}
	ce[2] = student{2, "wang", 23}
	fmt.Println(ce)
	delete(ce, 2)
	fmt.Println(ce)
}

func testJSON() {
	c := &Class{
		Title:    "101",
		Students: make([]*Student_, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student_{
			Name:   fmt.Sprintf("stu%02d", i),
			gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	//JSON反序列化：JSON格式的字符串-->结构体
	str := data
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}

func newPerson(name, city string, age int8) *Person {
	return &Person{
		name: name,
		city: city,
		age:  age,
	}
}
func mapTest3() {
	m := make(map[string]*Student)
	stus := []Student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}
	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
func mapTest2() {
	b := make(map[string]int, 1)
	b["测试"] = 100
	b["测试1"] = 101
	fmt.Println(b)
	fmt.Println("length:", len(b))
	fmt.Printf("type:%T", b)
}
func mapTest1() {
	p1 := Person{}
	p1.name = "pprof.cn"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)  //p1={pprof.cn 北京 18}
	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"pprof.cn", city:"北京", age:18}

	p2 := new(Person)
	p2.name = "pprof.cn"
	p2.city = "北京"
	p2.age = 19
	fmt.Printf("%T\n", p2)
	fmt.Printf("p2=%#v\n", p2)
}
func mapTemp() {
	//直接创建初始化一个map
	//var mapInit = map[string]string{"xiaoli": "湖南", "xiaoliu": "天津"}
	//声明一个map类型变量,
	//map的key的类型是string，value的类型是string
	var mapTemp map[string]string
	//使用make函数初始化这个变量,并指定大小(也可以不指定)
	mapTemp = make(map[string]string, 10)
	//存储key ，value
	mapTemp["xiaoming"] = "北京"
	mapTemp["xiaowang"] = "河北"
	//根据key获取value,
	//如果key存在，则ok是true，否则是flase
	//v1用来接收key对应的value,当ok是false时，v1是nil
	v1, ok := mapTemp["xiaoming"]
	fmt.Println(ok, v1)
	//当key=xiaowang存在时打印value
	if v2, ok := mapTemp["xiaowang"]; ok {
		fmt.Println(v2)
	}
	//遍历map,打印key和value
	for k, v := range mapTemp {
		fmt.Println(k, v)
	}
	fmt.Println(len(mapTemp))
	//删除map中的key
	delete(mapTemp, "xiaoming")
	//获取map的大小
	l := len(mapTemp)
	fmt.Println(l)
}
func mapGo() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}
func sliceMap() {
	sliceMap := make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Printf("type:%T\n", sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	fmt.Println(sliceMap[key])
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
func scoreMap() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

/*
	func makeslice(et *_type, len, cap int) slice {
		// 根据切片的数据类型，获取切片的最大容量
		maxElements := maxSliceCap(et.size)
		// 比较切片的长度，长度值域应该在[0,maxElements]之间
		if len < 0 || uintptr(len) > maxElements {
			panic(errorString("makeslice: len out of range"))
		}
		// 比较切片的容量，容量值域应该在[len,maxElements]之间
		if cap < len || uintptr(cap) > maxElements {
			panic(errorString("makeslice: cap out of range"))
		}
		// 根据切片的容量申请内存
		p := mallocgc(et.size*uintptr(cap), et, true)
		// 返回申请好内存的切片的首地址
		return slice{p, len, cap}
	}
*/
func sliceTest1() {
	slice := make([]byte, 3)
	n := copy(slice, "abcdef")
	fmt.Println(n, slice)
}
func arrTest2() {
	arrayA := []int{100, 200}
	testArrayPoint(&arrayA) // 1.传数组指针
	arrayB := arrayA[:2]
	testArrayPoint(&arrayB) // 2.传切片
	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
}
func arrTest1() {
	arrayA := [2]int{100, 200}
	var arrayB [2]int

	arrayB = arrayA

	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
	fmt.Printf("arrayB : %p , %v\n", &arrayB, arrayB)

	testArray(arrayA)
}
func testArrayPoint(x *[]int) {
	fmt.Printf("func Array : %p , %v\n", x, *x)
	(*x)[1] += 100
}
func testArray(x [2]int) {
	fmt.Printf("func Array : %p , %v\n", &x, x)
}
