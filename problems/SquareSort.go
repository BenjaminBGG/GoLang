/*SquareSort - square values in sorted array of integers; return sorted array */

package main

import (
	"fmt"
	"math"
)

func main() {
	var arr = []int{-11, -9, -6, -4, 0, 4, 7, 54}
	var arr2 = []int{}
	var arr3 = []int{0}
	var arr4 = []int{-99}
	fmt.Println(SquareSort(arr))
	fmt.Println("[0 16 16 36 49 81 121 2916] Expected")
	fmt.Println(SquareSort(arr2))
	fmt.Println("[-1] Expected")
	fmt.Println(SquareSort(arr3))
	fmt.Println("[0] Expected")
	fmt.Println(SquareSort(arr4))
	fmt.Println("[9801] Expected")
}

/*SquareSort - takes sorted array of integers as input, returns squared values in sorted array of integers
params:
	arr []int - sorted array of integers
returns:
	ret []int - sorted array of integers
*/
func SquareSort(arr []int) []int {
	var left, right int = 0, 1
	var imin, imax bool = false, false
	ret := []int{}
	if len(arr) < 1 {
		ret = append(ret, -1)
		return ret
	}
	if len(arr) == 1 {
		ret = append(ret, int(math.Pow(float64(arr[0]), 2)))
		return ret
	}
	for arr[right] < 0 && right < len(arr)-1 {
		right++
		left++
	}
	for !(imin) || !(imax) {
		if !(imin) && (imax || (int(math.Pow(float64(arr[left]), 2)) <= int(math.Pow(float64(arr[right]), 2)))) {
			ret = append(ret, int(math.Pow(float64(arr[left]), 2)))
			if left > 0 {
				left--
			} else {
				imin = true
			}
		} else if !(imax) && right <= len(arr)-1 {

			ret = append(ret, int(math.Pow(float64(arr[right]), 2)))
			if right < len(arr)-1 {
				right++
			} else {
				imax = true
			}
		}
	}
	return ret
}
