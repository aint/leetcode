package main

func sortArrayByParityII(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if (i % 2 == 0 && nums[i] % 2 == 0) || (i % 2 != 0 && nums[i] % 2 != 0) {
			continue
		}

		if i % 2 == 0 {
			// search even & swap
			for j := i; j < len(nums); j++ {
				if nums[j] % 2 == 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		} else {
			// search odd & swap
			for j := i; j < len(nums); j++ {
				if nums[j] % 2 != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}

	return nums
}
