package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// var wg sync.WaitGroup
	// Test0(&wg)
	// fmt.Println("main end")

	// Test1()
	// fmt.Println("main end")

	// c := &Counter{0}
	// Test2(c)
	// fmt.Println("--调用Test2后的c:", c)
	// fmt.Println("--调用Test2后的值:", c.GetCount())
	// fmt.Println("main end")

	// c := &CounterSync{}
	// Test3(c)
	// fmt.Println("main end")

	// c := &CounterSync{}
	// Test4(c)
	// fmt.Println("main end")

	//获取结构体指针
	// c := &CounterSync{}
	// Test5(c)
	// fmt.Println("main end")

	// c := &AtomicCounter{}
	// Test6(c)

	Test7()
}

func Test0(wg *sync.WaitGroup) {
	var mu sync.Mutex
	num := 0

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			defer wg.Done()
			defer mu.Unlock()
			for i := 0; i < 100; i++ {
				num += 1
			}
		}()
	}
	wg.Wait()
	fmt.Println("累加后的结果num：", num)
}

func Test1() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	num := 0

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			defer wg.Done()
			defer mu.Unlock()
			for i := 0; i < 100; i++ {
				num += 1
			}
		}()
	}
	wg.Wait()
	fmt.Println("累加后的结果num：", num)
}

type Counter struct {
	Count int
}

func (c *Counter) Increment() {
	c.Count++
}
func (c *Counter) GetCount() int {
	return c.Count
}

func Test2(c *Counter) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			//任务完成后通知
			defer wg.Done()
			// 获取锁
			mu.Lock()
			//函数返回时释放锁
			defer mu.Unlock()
			for j := 0; j < 100; j++ {
				c.Increment()
			}
		}()
	}
	//等待所有goroutine执行完毕
	wg.Wait()
}

type CounterSync struct {
	count int
	//互斥锁，用于保护共享数据
	mu sync.Mutex
}

func (c *CounterSync) Increment() {
	//获取锁
	c.mu.Lock()
	//函数返回时释放锁
	defer c.mu.Unlock()
	c.count++
}
func (c *CounterSync) GetCount() int {
	//获取锁
	c.mu.Lock()
	//函数返回时释放锁
	defer c.mu.Unlock()
	return c.count
}

func Test3(c *CounterSync) {
	for i := 0; i < 10000; i++ {
		c.Increment()
	}
	fmt.Println("任务执行完成后的结果count：", c.GetCount())
}

func Test4(c *CounterSync) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			//任务完成后通知
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				c.Increment()
			}
			fmt.Println("第", i, "个执行任务")
		}(i)
	}
	wg.Wait()
	fmt.Println("任务执行完成后的结果count：", c.GetCount())
}

func Test5(c *CounterSync) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			//任务完成后通知
			defer wg.Done()
			for j := 0; j < 100; j++ {
				c.Increment()
			}
			fmt.Println("第", i, "个执行任务")
		}(i)
	}
	wg.Wait()
	fmt.Println("任务执行完成后的结果count：", c.GetCount())
}

// 原子操作的无锁计数器
type AtomicCounter struct {
	count int64
}

// 结构体方法
func (a *AtomicCounter) Increment() {
	//atomic方法需要传递指针
	atomic.AddInt64(&a.count, 1)
}
func (a *AtomicCounter) GetCount() int64 {
	////atomic方法需要传递指针
	return atomic.LoadInt64(&a.count)
}

func Test6(c *AtomicCounter) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			//任务完成后通知
			defer wg.Done()
			for j := 0; j < 100; j++ {
				c.Increment()
			}
			fmt.Println("第", i, "个执行任务")
		}(i)
	}
	wg.Wait()
	fmt.Println("任务执行完成后的结果count：", c.GetCount())
}

func Test7() {
	var count int64
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		//添加任务
		wg.Add(1)
		go func() {
			// 任务完成后通知
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				//atomic方法需要传递指针
				atomic.AddInt64(&count, 1)
			}
		}()
	}
	//等待所有 goroutine 完成
	wg.Wait()
	fmt.Println("计算结束后count:", count)
}
