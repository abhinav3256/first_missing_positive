package main

import (
	"fmt"
)

func main() {

	nums := []int{3, 4, -1, 1}
	j := 1
	for i := 0; i < len(nums); {
		//fmt.Println("i :", i)
		if nums[i] == j {
			//fmt.Println(nums[i], i, j)
			i = 0
			j = j + 1

			//fmt.Println(nums[i], i, j)
		} else if i == len(nums)-1 {
			fmt.Println("j", j, i)
			break

		} else {
			i++
		}

	}

}
