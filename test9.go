package main

import "fmt"

func main()  {
	/*sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum) //45*/

	sum := 0
	for {
		sum++
		if sum > 100 {
			break
		}
	}
	fmt.Println("sum = ", sum) //101

}
