package task

import "fmt"

// 供外部包调用，首字母大写
func Test1(i int) int {
	//函数传值，传递的是原始参数的副本，方法中的修改不影响原始数据
	i += 10
	fmt.Println("方法中修改入参后：", i)
	return i
}

func IsValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	//定义map，指定符号关系
	mp := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	//初始化一个空的切片
	stack := []string{}
	for i, v := range s {
		t := string(v)
		fmt.Println("字符串str[", i, "]=：", t)
		fmt.Println("stack长度", len(stack))
		if len(stack) == 0 {
			stack = append(stack, t)
		} else {
			_, exist := mp[t]
			if exist {
				stack = append(stack, t)
			} else {
				//取出最新添加进去的元素
				sstr := stack[len(stack)-1]
				//切片元素减一
				stack = stack[:len(stack)-1]
				m2 := mp[sstr]
				if string(v) != m2 {
					return false
				}
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
