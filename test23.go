package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	//Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可

	go say("world")
	say("hello")

	//通道（channel）是用来传递数据的一个数据结构
	//ch <- v    // 把 v 发送到通道 ch
	//v := <-ch  // 从 ch 接收数据 并把值赋给 v

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	fmt.Println(s[:len(s)/2]) //[7 2 8]

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y) //-5 17 12

	// 这里我们定义了一个可以存储整数类型的带缓冲通道
	// 缓冲区大小为2
	ch1 := make(chan int, 2)

	// 因为 ch1 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	ch1 <- 1
	ch1 <- 2

	// 获取这两个数据
	fmt.Println(<-ch1) //1
	fmt.Println(<-ch1) //2

	cc := make(chan int, 10)
	go fibonacci(cap(cc), cc)
	// range 函数遍历每个从通道接收到的数据，因为 cc 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 cc 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range cc {
		fmt.Println(i)
	}

}
