package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// r := &(Rectangle{10, 5})
	// r.Area()
	// r.Perimeter()
	// c := &(Circle{10})
	// c.Area()
	// c.Perimeter()

	// p := Person{"Golang", 20}
	// e := &Employee{1, p}
	// e.PrintInfo()

	// c := make(chan int)
	// Ch1(c)
	// time.Sleep(5 * time.Second)

	// c := make(chan int, 200)
	// c := make(chan int, 2)
	// Ch2(c)
	// fmt.Println("main end")

	// c := make(chan int)
	// Ch3(c)
	// time.Sleep(15 * time.Second)
	// fmt.Println("main end")

	c := make(chan int, 100)
	Ch4(c)
	fmt.Println("main end")
}

/**
 * 面向对象
 */

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	Length int
	Width  int
}

func (r *Rectangle) Area() {
	t := r.Length * r.Width
	fmt.Println("Rectangle is Area：", t)
}
func (r *Rectangle) Perimeter() {
	t := (r.Length + r.Width) * 2
	fmt.Println("Rectangle is Perimeter：", t)
}

type Circle struct {
	Radius int
}

func (r *Circle) Area() {
	t := float64(r.Radius) * float64(r.Radius) * 3.14
	fmt.Println("Circle is Area：", t)
}
func (r *Circle) Perimeter() {
	t := 2 * float64(r.Radius) * 3.14
	fmt.Println("Circle is Perimeter：", t)
}

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	EmployeeID int
	Person
}

func (e *Employee) PrintInfo() {
	fmt.Println(e)
	fmt.Printf("员工id：%d，姓名：%s, 年龄：%d\n", e.EmployeeID, e.Name, e.Age)
}

func Ch1(t chan int) {
	//生产者：生成数字并发送到通道
	go func() {
		for i := 1; i <= 10; i++ {
			t <- i
			fmt.Println("Goroutine向channel写数据：", i)
			time.Sleep(100 * time.Millisecond)
		}
		close(t)
		fmt.Println("生产者完成")
	}()

	// 消费者：从通道接收并打印数字
	go func() {
		for v := range t {
			fmt.Println("Goroutine读取channel数据：", v)
		}
		fmt.Println("消费者完成")
	}()
}

/**
 * sync.WaitGroup同步控制，主进程会等待goroutine 返回后才执行结束
 */
func Ch2(t chan int) {
	var wg sync.WaitGroup
	//生产者：生成数字并发送到通道
	//添加任务
	wg.Add(1)
	go func() {
		// 任务完成后通知
		defer wg.Done()
		// 任务完成后关闭channel
		defer close(t)
		for i := 1; i <= 10; i++ {
			t <- i
			fmt.Println("Goroutine向channel写数据：", i)
		}
		fmt.Println("生产者完成")
	}()

	// 消费者：从通道接收并打印数字
	//添加任务
	wg.Add(1)
	go func() {
		// 任务完成后通知
		defer wg.Done()
		for v := range t {
			fmt.Println("Goroutine读取channel数据：", v)
		}
		fmt.Println("消费者完成")
	}()
	wg.Wait()
}

func Ch3(t chan int) {
	//生产者：生成数字并发送到通道
	go func() {
		for i := 1; i <= 100; i++ {
			t <- i
			fmt.Println("Goroutine向channel写数据：", i)
			time.Sleep(100 * time.Millisecond)
		}
		close(t)
		fmt.Println("生产者完成")
	}()

	// 消费者：从通道接收并打印数字
	go func() {
		for v := range t {
			fmt.Println("Goroutine读取channel数据：", v)
		}
		fmt.Println("消费者完成")
	}()
}

func Ch4(t chan int) {
	var wg sync.WaitGroup
	//生产者：生成数字并发送到通道
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(t)
		for i := 1; i <= 100; i++ {
			t <- i
			fmt.Println("Goroutine向channel写数据：", i)
		}
		fmt.Println("生产者完成")
	}()

	// 消费者：从通道接收并打印数字
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range t {
			fmt.Println("Goroutine读取channel数据：", v)
		}
		fmt.Println("消费者完成")
	}()

	wg.Wait()
}
