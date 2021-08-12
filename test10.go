package main

import "fmt"

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {

	var num int
	num = max(2, 3)
	fmt.Println("num = ", num) //3

	a, b := swap("Google", "Runoob")
	fmt.Println(a, b) //Runoob  Google
}
