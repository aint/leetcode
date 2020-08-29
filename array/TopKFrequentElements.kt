class Solution {
    fun topKFrequent(nums: IntArray, k: Int): IntArray {
        val map = hashMapOf<Int, Int>()
        var topInt = nums[0]
        var maxCount = 0

        for (i in nums.indices) {
            var count = map[nums[i]]
            if (count == null) {
                map[nums[i]] = 1
            } else {
                count += 1
                map[nums[i]] = count
                if (count > maxCount) {
                    maxCount = count
                    topInt = nums[i]
                }
            }
        }

        val result = IntArray(k)
        result[0] = topInt

        map.remove(topInt)
        maxCount = 0

        for (i in 1 until k) {
            for (entry in map.entries) {
                if (entry.value > maxCount) {
                    maxCount = entry.value
                    topInt = entry.key
                }
            }
            result[i] = topInt
            map.remove(topInt)
            maxCount = 0
        }

        return result
    }
}

fun main() {
    val ints = Solution().topKFrequent(intArrayOf(1, 1, 1, 2, 2, 3), 2)
    println("RESULT = " + ints.contentToString())
}