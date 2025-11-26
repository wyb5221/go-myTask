package main

import (
	"fmt"
	task1 "go-myTask/task1"
	task2 "go-myTask/task2"
)

func main() {

	i := 5
	task1.Test1(i)
	fmt.Println("调用test1后的值i：", i)
	task2.Test2(&i)
	fmt.Println("调用test2后的值i：", i)
}
