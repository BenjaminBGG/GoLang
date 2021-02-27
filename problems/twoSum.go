/*
two_sum in GoLang -- find the two numbers in a sorted array which add to a third number
params:
arr[] - sorted array of integers
num - target integer
*/

package main

import (
	"fmt"
)

func main() {
	var arr = []int{-11, -9, -6, -4, 0, 4, 7, 54}
	var num1 = 0
	var num2 = 48
	var num3 = -15
	var num4 = 77
	fmt.Println(TwoSum(arr, num1))
	fmt.Printf("Should be [3 5]\n")
	fmt.Println(TwoSum(arr, num2))
	fmt.Printf("Should be [2 7]\n")
	fmt.Println(TwoSum(arr, num3))
	fmt.Printf("Should be [0 3]\n")
	fmt.Println(TwoSum(arr, num4))
	fmt.Printf("Should be [-1 -1]")
}

// TwoSum -- Returns array position of two numbers which add to input
func TwoSum(arr []int, num int) []int {
	var left, right int = 0, len(arr) - 1
	ret := []int{}
	for left <= right {
		if arr[left]+arr[right] < num {
			left++
		}
		if arr[left]+arr[right] > num {
			right--
		}
		if arr[left]+arr[right] == num {
			ret = append(ret, left)
			ret = append(ret, right)
			return ret
		}
	}
	ret = append(ret, -1)
	ret = append(ret, -1)
	return ret
}
