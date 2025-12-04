package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// i := 5
	// Test1(i)
	// fmt.Println("调用Test1后的值i：", i)
	// Test2(&i)
	// fmt.Println("调用Test2后的值i：", i)

	// i := []int{1, 5, 9}
	// Test3(i)
	// fmt.Println("调用Test3后的值i：", i)
	// Test4(&i)
	// fmt.Println("调用Test4后的值i：", i)

	// Test5()
	// time.Sleep(2 * time.Second)

	Test6()
}

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

func Test3(its []int) []int {
	fmt.Println("初始化数据：", its)
	for i, v := range its {
		a := &v
		fmt.Println("切片中第[", i, "]个元素：", *a)
		*a *= 2
		fmt.Println("修改后数据：", *a)
	}
	fmt.Println("修改后数据：", its)
	return its
}

func Test4(its *[]int) {
	fmt.Println("初始化数据：", its)
	//循环中的v是切片中元素的副本
	for i, v := range *its {
		fmt.Println("切片中第[", i, "]个元素：", v)
		//直接修改切片中的数据
		(*its)[i] = v * 2
	}
}

func Test5() {
	go func() {
		fmt.Print("第一个Goroutine打印：")
		for i := 1; i < 10; i += 2 {
			fmt.Print(i, " , ")
		}
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("")
	go func() {
		fmt.Print("第二个Goroutine打印：")
		for i := 2; i <= 10; i += 2 {
			fmt.Print(i, " , ")
		}
	}()
}

func Test6() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Print("第一个Goroutine打印：")
		for i := 1; i < 10; i += 2 {
			fmt.Print(i, " , ")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("")
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Print("第二个Goroutine打印：")
		for i := 2; i <= 10; i += 2 {
			fmt.Print(i, " , ")
		}
	}()

	wg.Wait()
}
