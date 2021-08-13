package main

import "fmt"

/* 声明全局变量 */
var g int

/* 声明全局变量 */
var gg int = 20

/* 声明全局变量 */
var aa int = 20

/* 函数定义-两数相加 */
func sum(aa, ba int) int {
	fmt.Printf("sum() 函数中 aa = %d\n", aa)
	fmt.Printf("sum() 函数中 ba = %d\n", ba)

	return aa + ba
}

func main() {
	/* 声明局部变量 */
	var a, b, c int

	/* 初始化参数 */
	a = 10
	b = 20
	c = a + b

	fmt.Printf("结果： a = %d, b = %d and c = %d\n", a, b, c) //10 20 30

	g = a + b

	fmt.Printf("结果： a = %d, b = %d and g = %d\n", a, b, g) //10 20 30

	/* 声明局部变量 */
	var gg int = 10

	fmt.Printf("结果： gg = %d\n", gg) //10

	/* main 函数中声明局部变量 */
	var aa int = 10
	var ba int = 20
	var ca int = 0

	fmt.Printf("main()函数中 aa = %d\n", aa) //10
	ca = sum(aa, ba) // 10 20
	fmt.Printf("main()函数中 ca = %d\n", ca) //30

	//不同类型的局部和全局变量默认值为 int,float32 为0     pointer 为nil
}
