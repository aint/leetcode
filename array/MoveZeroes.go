package main

func removeElement(nums []int, val int) int {
	count := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			moveElementToEnd(nums, i)
			count++
		}
    }
	return len(nums) - count;
}

func moveElementToEnd(nums []int, index int) {
	for i := index; i < len(nums) - 1; i++ {
		val := nums[i]
		nums[i] = nums[i + 1]
		nums[i + 1] = val
	}
}