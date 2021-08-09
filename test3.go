package main

import "fmt"

//Go 程序可以由多个标记组成，可以是关键字，标识符，常量，字符串，符号
//以下是无效的标识符:1.1ab（以数字开头）  2.case（Go 语言的关键字）  3.a+b（运算符是不允许的）
func main() {
	fmt.Println("Hello,ni hao!")
	//Go 语言的字符串可以通过 + 实现
	fmt.Println("Google" + "Runoob") //GoogleRunoob

	var age int
	fmt.Println("age = ", age) //age =  0

	//格式化字符串
	// %d 表示整型数字，%s 表示字符串
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url) //Code=123&endDate=2020-12-31
}
