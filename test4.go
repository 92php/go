package main

import "fmt"

/**
 * 一个结构体就是一个字段的集合
 */
type Vertex struct {
	X int
	Y float64
}

func add(a, b int) int {
	return a + b
}

type FuncType func(int, int) int

type Humaner interface {
	sayHi()
}

func sayHi() {
	fmt.Println("hi")
}

func main() {
	//布尔型
	var b bool = true
	fmt.Println("b = ", b)

	//数字类型  整型 int 和浮点型 float32、float64
	var year int = 2021
	fmt.Println("year = ", year)

	var f1 float32 = 3.14
	fmt.Printf("f1的type is %T\n", f1)
	fmt.Printf("f1 = %3.2f\n", f1)
	// 自动推导类型
	f2 := 3.45
	fmt.Printf("f2的type is %T\n", f2) // float64
	fmt.Printf("f2 = %3.2f\n", f2)

	//字符串类型 string
	var str string
	str = "aaa"
	fmt.Println("str = ", str)

	//派生类型  1.指针类型（Pointer） 2.数组类型 3.结构化类型(struct) 4.Channel 类型  5.函数类型 6.切片类型 7.接口类型（interface） 8.Map 类型
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */
	ip = &a        /* 指针变量的存储地址 */
	fmt.Printf("a 变量的地址是: %x\n", &a)
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip) /* 指针变量的存储地址 */
	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	var c [5]int           //定义数组 5个整数型
	fmt.Println("emp:", c) //整数型值为0 emp: [0 0 0 0 0]

	c[4] = 100                //修改数组
	fmt.Println("set:", c)    //set: [0 0 0 0 100]
	fmt.Println("get:", c[4]) //输出单个值 100

	fmt.Println("len:", len(c)) //数组长度 5

	d := [5]int{45, 32, 12, 42, 55} //定义数组 定义数组的值
	fmt.Println("dcl:", d)

	for z, y := range d { //遍历数组
		fmt.Println("ran:", z, y)
	}

	fmt.Println(Vertex{1, 2.31231})

	var c1 chan int
	fmt.Println("c1 = ", c1)

	var result int
	var fTest FuncType
	fTest = add
	result = fTest(10, 20)
	fmt.Println("result = ", result)

	//定义一个切片和定义一个数组的语法相似，唯一的不同是不需要定义切片长度
	letters := []string{"a", "b", "c", "d"}
	fmt.Println("letters = ", letters)
	var s []byte
	s = make([]byte, 5, 5)
	fmt.Println("s = ", s)

	var i Humaner
	fmt.Println("i = ", i)

	//通过make函数 创建一个映射，键类型是string ，值是int
	list := make(map[string]int)
	list["test"] = 1
	fmt.Println(list["test"])
	//创建一个映射，键值都是string  使用字面量
	data := map[string]string{"top": "is top", "bottom": "is bottom"}
	fmt.Println(data["top"])

	//数字类型 uint8无符号 8 位整型
	//uint16 无符号 16 位整型
	//uint32 无符号 32 位整型
	//uint64 无符号 64 位整型
	//int8  有符号 8 位整型
	//int16 有符号 16 位整型
	//int32 有符号 32 位整型
	//int64 有符号 64 位整型

	//浮点型 float32 32位浮点型数
	//float64  64位浮点型数
	//complex64  32 位实数和虚数
	//complex128  64 位实数和虚数

	//其他数字类型
	//byte 类似 uint8
	//rune 类似 int32
	//uint 32 或 64 位
	//int 与 uint 一样大小
	//uintptr  无符号整型，用于存放一个指针

}
