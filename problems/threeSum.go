/*threeSum -- Given a sorted array of integers, find three numbers which add to a given int */
package main

import "fmt"

func main() {
	var arr = []int{-11, -9, -6, -4, 0, 4, 7, 54}
	var num1 = 0
	var num2 = 45
	var num3 = -8
	var num4 = 77
	fmt.Println(ThreeSum(arr, num1))
	fmt.Printf("Should be [0 5 6]\n")
	fmt.Println(ThreeSum(arr, num2))
	fmt.Printf("Should be [1 4 7]\n")
	fmt.Println(ThreeSum(arr, num3))
	fmt.Printf("Should be [0 3 6]\n")
	fmt.Println(ThreeSum(arr, num4))
	fmt.Printf("Should be [-1 -1 -1]")
}

/*
ThreeSum -- Given a sorted array of integers, find three numbers which add to a given int
params:
	arr []int: 	Sorted array of integers
	num int: 	target integer
*/
func ThreeSum(arr []int, num int) []int {
	ret := []int{}
	var i int = 0
	for i = 0; i < len(arr)-1; i++ {
		var left, right int = i + 1, len(arr) - 1
		for left < right {
			if arr[i]+arr[left]+arr[right] < num {
				left++
			}
			if arr[i]+arr[left]+arr[right] > num {
				right--
			}
			if arr[i]+arr[left]+arr[right] == num {
				ret = append(ret, i)
				ret = append(ret, left)
				ret = append(ret, right)
				return ret
			}
		}
	}
	ret = append(ret, -1)
	ret = append(ret, -1)
	ret = append(ret, -1)
	return ret
}
