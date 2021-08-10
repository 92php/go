package main

import (
	"fmt"
	"unsafe"
)

const (
	aa = "abc"
	ba = len(aa)
	ca = unsafe.Sizeof(aa)
)

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

func main() {
	//常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d\n", area)
	println()

	println(a, b, c) // 1 false str
	println()

	println(aa, ba, ca) //abc 3 16
	println()

	const (
		ax = iota //0
		bx        //1
		cx        //2
		dx = "ha" //独立值，iota += 1
		ex        //"ha"   iota += 1
		fx = 100  //iota +=1
		gx        //100  iota +=1
		hx = iota //7,恢复计数
		ix        //8
	)
	fmt.Println(ax, bx, cx, dx, ex, fx, gx, hx, ix) //0 1 2 ha ha 100 100 7 8

	fmt.Println("i=", i) //i=1：左移 0 位，不变仍为 1
	fmt.Println("j=", j) //j=3：左移 1 位，变为二进制 110，即 6
	fmt.Println("k=", k) //k=3：左移 2 位，变为二进制 1100，即 12
	fmt.Println("l=", l) //l=3：左移 3 位，变为二进制 11000，即 24

}
