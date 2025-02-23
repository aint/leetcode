fun removeDuplicates(nums: IntArray): Int {
    // Use a two-pointer approach here. One, that would keep track of the current element
    // in the original array and another one for just the unique elements.
    // Essentially, once an element is encountered, you simply need to bypass its duplicates
    // and move on to the next unique element.
    var k = 1
    var prev = nums[0]
    for (i in nums.indices) {
        if (i == 0 || nums[i] == prev) {
            continue
        }
        nums[k] = nums[i]
        prev = nums[i]
        k++
    }
    // println(nums.joinToString(", "))
    return k
}

fun main() {
    removeDuplicates(intArrayOf(0, 0, 1, 1, 1, 2, 2, 3, 3, 4))
}
