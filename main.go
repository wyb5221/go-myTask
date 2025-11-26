package main

import (
	"fmt"
	_ "fmt"
	_ "go-myTask/task1"
	_ "go-myTask/task2"
	"go-myTask/task3"
	"sync"
)

func main() {

	// i := 5
	// f := task1.IsValid("([])")
	// fmt.Println("调用IsValid后的值i：", f)
	// task2.Test2(&i)
	// fmt.Println("调用test2后的值i：", i)
	// its := []int{1, 2, 3}
	// task2.Test4(&its)
	// fmt.Println("调用test4后的值：", its)
	// task2.Test5()
	// time.Sleep(3 * time.Second)
	// fmt.Println("main 退出")

	// r := &(task3.Rectangle{10, 5})
	// r.Area()
	// r.Perimeter()
	// c := &(task3.Circle{10})
	// c.Area()
	// c.Perimeter()

	// p := task3.Person{"Golang", 20}
	// e := &task3.Employee{1, p}
	// e.PrintInfo()

	// c := make(chan int)
	// task3.Ch1(c)
	// c := make(chan int, 200)
	// c := make(chan int, 2)
	// var w sync.WaitGroup
	// task3.Ch2(c, w)
	// fmt.Println("main end")
	// c := make(chan int)
	// task3.Ch3(c)
	// time.Sleep(20 * time.Second)
	// fmt.Println("main end")
	c := make(chan int, 100)
	var wg sync.WaitGroup
	task3.Ch4(c, wg)
	fmt.Println("main end")
}
