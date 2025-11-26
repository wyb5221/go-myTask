package main

import (
	"fmt"
	task1 "go-myTask/task1"
	task2 "go-myTask/task2"
)

func main() {

	i := 5
	f := task1.IsValid("([])")
	fmt.Println("调用IsValid后的值i：", f)
	task2.Test2(&i)
	fmt.Println("调用test2后的值i：", i)
}
