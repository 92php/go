package main

import "fmt"

func main() {
	/*//声明数组
	var balance [10]float32
	fmt.Println("balance = ", balance)

	//初始化数组
	var balance1 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println("balance1 = ", balance1)

	balance2 := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println("balance2 = ", balance2)

	balance3 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println("balance3 = ", balance3)

	//  将索引为 1 和 3 的元素初始化
	balance4 := [5]float32{1:2.0,3:7.0}
	fmt.Println("balance4 = ", balance4)*/

	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}


	var ii, ji, k int
	// 声明数组的同时快速初始化数组
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	/* 输出数组元素 */
	for ii = 0; ii < 5; ii++ {
		fmt.Printf("balance[%d] = %f\n", ii, balance[ii])
	}

	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	/* 输出每个数组元素的值 */
	for ji = 0; ji < 5; ji++ {
		fmt.Printf("balance2[%d] = %f\n", ji, balance2[ji])
	}

	//  将索引为 1 和 3 的元素初始化
	balance3 := [5]float32{1: 2.0, 3: 7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}

	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 appped() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)

	// Step 3: 显示两行数据
	fmt.Println("Row 1")
	fmt.Println(values[0])
	fmt.Println("Row 2")
	fmt.Println(values[1])

	// Step 4: 访问第一个元素
	fmt.Println("第一个元素为：")
	fmt.Println(values[0][0])

	// 创建二维数组
	sites := [2][2]string{}

	// 向二维数组添加元素
	sites[0][0] = "Google"
	sites[0][1] = "Runoob"
	sites[1][0] = "Taobao"
	sites[1][1] = "Weibo"

	// 显示结果
	fmt.Println(sites)



	/* 数组 - 5 行 2 列*/
	var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
	var ij, jj int

	/* 输出数组元素 */
	for  ij = 0; ij < 5; ij++ {
		for jj = 0; jj < 2; jj++ {
			fmt.Printf("a[%d][%d] = %d\n", ij,jj, a[ij][jj] )
		}
	}


	// 创建空的二维数组
	animals := [][]string{}

	// 创建三一维数组，各数组长度不同
	row11 := []string{"fish", "shark", "eel"}
	row21 := []string{"bird"}
	row31 := []string{"lizard", "salamander"}

	// 使用 append() 函数将一维数组添加到二维数组中
	animals = append(animals, row11)
	animals = append(animals, row21)
	animals = append(animals, row31)

	// 循环输出
	for m := range animals {
		fmt.Printf("Row: %v\n", m)
		fmt.Println(animals[m])
	}


	/* 数组长度为 5 */
	var  balance55 = [5]int {1000, 2, 3, 17, 50}
	var avg float32

	/* 数组作为参数传递给函数 */
	avg = getAverage( balance55, 5 )

	/* 输出返回的平均值 */
	fmt.Printf( "平均值为: %f ", avg )

}

func getAverage(arr [5]int, size int) float32 {
	var i,sum int
	var avg float32

	for i = 0; i < size;i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}
