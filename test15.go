package main

import "fmt"

func main() {
	//Go 语言切片是对数组的抽象
	//与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大

	//直接初始化切片，[] 表示是切片类型，{1,2,3} 初始化值依次是 1,2,3，其 cap=len=3
	s := []int{1, 2, 3}

	fmt.Println("s = ", s) //s =  [1 2 3]

	var numbers = make([]int, 3, 5)

	printSlice(numbers) //len=3 cap=5 slice=[0 0 0]

	//空(nil)切片
	var numbers1 []int

	printSlice(numbers1) //len=0 cap=0 slice=[]

	if numbers1 == nil {
		fmt.Printf("切片是空的")
	}

	//切片截取
	/* 创建切片 */
	numbers2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers2) //len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8]

	/* 打印原始切片 */
	fmt.Println("numbers2 ==", numbers2) //numbers2 == [0 1 2 3 4 5 6 7 8]

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers2[1:4] ==", numbers2[1:4]) //numbers2[1:4] == [1 2 3]

	/* 默认下限为 0*/
	fmt.Println("numbers2[:3] ==", numbers2[:3]) //numbers2[:3] == [0 1 2]

	/* 默认上限为 len(s)*/
	fmt.Println("numbers2[4:] ==", numbers2[4:]) //numbers2[4:] == [4 5 6 7 8]

	numbers3 := make([]int, 0, 5)
	printSlice(numbers3) //len=0 cap=5 slice=[]

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	numbers4 := numbers3[:2]
	printSlice(numbers4) //len=2 cap=5 slice=[0 0]

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	numbers5 := numbers3[2:5]
	printSlice(numbers5) //len=3 cap=3 slice=[0 0 0]

	var numbers6 []int
	printSlice(numbers6) //len=0 cap=0 slice=[]

	/* 允许追加空切片 */
	numbers6 = append(numbers6, 0)
	printSlice(numbers6) //len=1 cap=1 slice=[0]

	/* 向切片添加一个元素 */
	numbers6 = append(numbers6, 1)
	printSlice(numbers6) //len=2 cap=2 slice=[0 1]

	/* 同时添加多个元素 */
	numbers6 = append(numbers6, 2, 3, 4)
	printSlice(numbers6) //len=5 cap=6 slice=[0 1 2 3 4]

	/* 创建切片 numbers7 是之前切片的两倍容量*/
	numbers7 := make([]int, len(numbers6), (cap(numbers6))*2)

	/* 拷贝 numbers6 的内容到 numbers7 */
	copy(numbers7, numbers6)
	printSlice(numbers7) //len=5 cap=12 slice=[0 1 2 3 4]

}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
