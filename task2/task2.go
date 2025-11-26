package task

import "fmt"

// 供外部包调用，首字母大写
func Test1(i int) int {
	//函数传值，传递的是原始参数的副本，方法中的修改不影响原始数据
	i += 10
	fmt.Println("方法中修改入参后：", i)
	return i
}

func Test2(i *int) int {
	//函数传指针，传递的是原始参数的引用地址，方法中的修改会同步修改原始数据
	*i += 10
	fmt.Println("方法中修改入参后：", *i)
	return *i
}

// func test3([]int) {
// 	for
// }
