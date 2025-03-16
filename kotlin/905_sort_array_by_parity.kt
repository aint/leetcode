fun sortArrayByParity(nums: IntArray): IntArray {
    // Use two pointers: one starting from the left and the other from the right.
    // If the left pointer points to an even number, move it to the right.
    // If the right pointer points to an odd number, move it to the left.
    // If left is odd and right is even, swap them.
    var left = 0
    var right = nums.size - 1
    while (left < right) {
        if (nums[left] % 2 == 0) {
            left++
        } else if (nums[right] % 2 == 0) {
            var t = nums[left]
            nums[left] = nums[right]
            nums[right] = t
        } else {
            right--
        }
    }
    return nums
}

fun main() {
    var res = sortArrayByParity(intArrayOf(3,1,2,4))
    println(res.joinToString(", "))
}
