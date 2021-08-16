package main

import "fmt"

func main() {
	//什么是指针 一个指针变量指向了一个值的内存地址
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a) //c000018090

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip) //c000018090

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip) //20

	var ptr *int

	fmt.Printf("ptr 的值为 : %x\n", ptr) //0

	/* ptr 不是空指针 */
	if ptr != nil {
		fmt.Printf("ptr1 的值为 : %x\n", ptr)
	}
	/* ptr 是空指针 */
	if ptr == nil {
		fmt.Printf("ptr2 的值为 : %x\n", ptr) //0
	}



}
