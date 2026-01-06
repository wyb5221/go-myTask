package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// a := test01(arr, 7)
	a := tes02(arr, 11, 0, len(arr)-1)
	fmt.Println("目标位置：", a)
}

// 二分查找
func test01(arr []int, target int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := (end-start)/2 + start
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}

func tes02(arr []int, target int, start int, end int) int {
	if start > end {
		return -1
	}
	mid := (end-start)/2 + start
	if arr[mid] == target {
		return mid
	}
	if arr[mid] > target {
		end = mid - 1
	} else {
		start = mid + 1
	}

	return tes02(arr, target, start, end)
}
