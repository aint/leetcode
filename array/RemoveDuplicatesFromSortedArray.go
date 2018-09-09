package leetcode

func removeDuplicates(nums []int) int {
	newLen := 0

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			for j := i; j < len(nums)-1; j++ {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
			newLen++
		}
	}
	return len(nums) - newLen
}
