fun sortArrayByParityII(nums: IntArray): IntArray {
    // Use two pointers. The first one starts from 0 and checks even indexes.
    // The second starts from the end and checks odd indexes.
    // Use two pointers:
    // - `l` starts at index 0 and moves forward by 2, ensuring even indices have even numbers.
    // - `r` starts at the last index and moves backward by 2, ensuring odd indices have odd numbers.
    var l = 0
    var r = nums.size - 1
    while (l < nums.size && r > 0) {
        if (nums[l] % 2 == 0)  {
            l += 2
        } else if (nums[r] % 2 != 0) {
            r -= 2
        } else {
            val tmp = nums[l]
            nums[l] = nums[r]
            nums[r] = tmp
            l += 2
            r -= 2
        }
    }
    return nums
}

fun main() {
    println("[4,5,2,7] = " + sortArrayByParityII(intArrayOf(4,2,5,7)).joinToString(", "))
    println("[4,3] = " + sortArrayByParityII(intArrayOf(3,4)).joinToString(", "))
    println("[4,1,2,1] = " + sortArrayByParityII(intArrayOf(4,1,2,1)).joinToString(", "))
    println("[2,3,4,1,4,3,0,1,0,3] = " + sortArrayByParityII(intArrayOf(2,3,1,1,4,0,0,4,3,3)).joinToString(", "))
}
