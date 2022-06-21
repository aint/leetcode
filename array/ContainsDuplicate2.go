package main

// 1. Store number index
// 2. Calc difference between current index and stored index for duplicates
// 3. Return false if no duplicate
func containsNearbyDuplicate(nums []int, k int) bool {
    m := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        index, exists := m[nums[i]]
        if exists {
            if math.Abs(float64(index - i)) <= float64(k) {
                return true
            }
        }
        m[nums[i]] = i        
    }
    return false
}
