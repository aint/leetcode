package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if len(nums) == 1 {
		if target > nums[0] {
			return 1
		} else {
			return 0
		}
	}
	idx := len(nums) / 2
	if nums[idx] < target {
		return idx + searchInsert(nums[idx:], target)
	} else if nums[idx] > target {
		return searchInsert(nums[:idx], target)
	} else {
		return idx
	}
}

func testSearchInsert() {
	fmt.Println("2 =", searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println("1 =", searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println("4 =", searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println("0 =", searchInsert([]int{1, 3, 5, 6}, 0))
}
