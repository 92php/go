package main

import "fmt"

var xx, yx int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	ax int
	bx bool
)

var cx, dx int = 1, 2
var ex, fx = 123, "hello"

func main() {
	//Go 语言变量名由字母、数字、下划线组成，其中首个字符不能为数字
	var a string = "hello"
	fmt.Println("a = ", a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// 声明一个变量并初始化
	var d = "RUNOOB"
	fmt.Println(d)

	// 没有初始化就为零值
	var e int
	fmt.Println(e) //0

	// bool 零值为 false
	var f bool
	fmt.Println(f) //false

	//以下几种类型为 nil
	//var aa *int  //nil
	//var aa []int   //[]
	//var aa map[string] int //map[]
	//var aa chan int  //nil
	//var aa func(string) int  //nil
	//var aa error // error 是接口  //nil
	//fmt.Println(aa)

	var ii int
	var ff float64
	var bb bool
	var ss string
	fmt.Printf("%v %v %v %q\n", ii, ff, bb, ss) // 0 0 false ""

	var mm = true
	fmt.Println(mm) //true

	fm := "Runoob" // var fm string = "Runoob"
	fmt.Println(fm)

	//这种不带声明格式的只能在函数体中出现
	gx, hx := 123, "hello"
	println(xx, yx, ax, bx, cx, dx, ex, fx, gx, hx) //0 0 0 false 1 2 123 hello 123 hello

	_, numb, strs := numbers() //只获取函数返回值的后两个
	fmt.Println(numb, strs) //2 str
}

//一个可以返回多个值的函数
func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}
