package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// i := 5
	// Test1(i)
	// fmt.Println("调用Test1后的值i：", i)

	// Test2(&i)
	// fmt.Println("调用Test2后的值i：", i)

	// f := IsValid("()[]{}")
	// fmt.Println("调用IsValid后的返回：", f)

	// s := []string{"abcd", "abab", "abcde"}
	// f := LongestCommonPrefix(s)
	// fmt.Println("调用LongestCommonPrefix后的返回：", f)

	// i := []int{9, 9, 9}
	// f := PlusOne(i)
	// fmt.Println("调用PlusOne后的返回：", f)

	// i := []int{1, 3, 1, 5, 3, 9}
	// f := RemoveDuplicates(i)
	// fmt.Println("调用RemoveDuplicates后的返回：", f)

	// i := []int{1, 5, 9, 10}
	// f := TwoSum(i, 11)
	// fmt.Println("调用TwoSum后的返回：", f)

	i := [][]int{{1, 3}, {5, 8}, {2, 5}, {11, 13}, {10, 15}}
	f := Merge(i)
	fmt.Println("调用Merge后的返回：", f)
}

// 供外部包调用，首字母大写
func Test1(i int) int {
	//函数传值，传递的是原始参数的副本，方法中的修改不影响原始数据
	i += 10
	fmt.Println("方法中修改入参后：", i)
	return i
}

// 供外部包调用，首字母大写
func Test2(i *int) int {
	//函数传值，传递的是原始参数的副本，方法中的修改不影响原始数据
	*i += 10
	fmt.Println("方法中修改入参后：", *i)
	return *i
}

/**
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
 */
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

/**
 * 编写一个函数来查找字符串数组中的最长公共前缀
 */
func LongestCommonPrefix(strs []string) string {
	result := ""
	res := []rune(result)

	//获取数组第一个元素
	s1 := strs[0]
outter:
	for i, v := range s1 {
		t := string(v)
		for j := 1; j < len(strs); j++ {
			tr := []rune(strs[j])
			if i >= len(tr) {
				break outter
			}
			tp := string(tr[i])
			if tp != t {
				break outter
			}
		}
		res = append(res, v)
	}
	return string(res)
}

/**
 * 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
 * 将大整数加 1，并返回结果的数字数组。
 */
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		//获取数组元素
		t := digits[i]
		if t != 9 {
			digits[i] += 1
			if i == len(digits)-1 {
				return digits
			} else {
				for j := i + 1; j < len(digits); j++ {
					digits[j] = 0
				}
				return digits
			}
		}
	}
	lg := len(digits) + 1
	arr := make([]int, lg)
	arr[0] = 1
	for a := 1; a < len(digits)+1; a++ {
		arr[a] = 0
	}
	return arr
}

/**
 * 删除有序数组中的重复项， 返回删除后数组的新长度
 */
func RemoveDuplicates(nums []int) int {
	fmt.Println("nums：", nums)
	mp := make(map[int]int)
	for _, v := range nums {
		_, exist := mp[v]
		if !exist {
			// fmt.Println("map：", len(mp))
			nums[len(mp)] = v
			mp[v] = v
		}
	}
	for i := len(mp); i < len(nums); i++ {
		nums[i] = -1
	}
	return len(mp)
}

/**
 * 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
 */
func TwoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		a := nums[i]
		_, exist := mp[target-a]
		if exist {
			return []int{mp[target-a], i}
		} else {
			mp[a] = i
		}
		// for j := i + 1; j < len(nums); j++ {
		//  b := nums[j]
		//  if (a + b) == target {
		//      return []int{i, j}
		//  }
		// }
	}
	return []int{}
}

/**
 * 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间
 */
func Merge(intervals [][]int) [][]int {
	fmt.Println("排序前intervals：", intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println("排序后intervals：", intervals)
	arr := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		s1 := intervals[i]
		s10 := s1[0]
		a1 := arr[len(arr)-1][1]
		//第一个数组的大值大于等于下一个数组的小值，说明有重合
		if a1 >= s10 {
			s11 := s1[1]
			arr[len(arr)-1][1] = int(math.Max(float64(s11), float64(a1)))
		} else {
			arr = append(arr, s1)
		}
	}
	fmt.Println("arr：", arr)

	return arr
}
