fun moveZeroes(nums: IntArray) {
    if (nums.size == 1) {
        return
    }

    // This function moves all zeros in the array to the end while maintaining the order of non-zero elements.
    // It uses a two-pointer approach:
    // - `i` iterates through the array.
    // - `k` tracks the position where the next non-zero element should be placed.
    // As we iterate through `nums`, whenever we find a non-zero element, we move it to index `k` and increment `k`.
    var k = 0
    for (i in nums.indices) {
        if (nums[i] != 0) {
            nums[k] = nums[i]
            k++
        }
    }

    // After moving all non-zero elements, fill the remaining positions with zeros.
    while (k < nums.size) {
        nums[k] = 0
        k++
    }
}

fun main() {
    val nums = intArrayOf(0, 1, 3, 0, 12)
    moveZeroes(nums)
    println("1, 3, 12, 0, 0 = " + nums.joinToString(", "))
}

// 0, 1, 3, 0, 12
//    i
// k
//
// 1, 1, 3, 0, 12
//       i
//    k
//
// 1, 3, 3, 0, 12
//          i
//       k
//
// 1, 3, 3, 0, 12
//             i
//       k
//
// 1, 3, 12, 0, 12