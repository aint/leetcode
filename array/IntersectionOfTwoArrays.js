var intersection = function(nums1, nums2) {
    let result = new Set([])

    for (let i = 0; i < nums1.length; i++) {
        let a = nums1[i];
        for (let j = 0; j < nums2.length; j++) {
            if (a == nums2[j]) {
                result.add(a)
            }
        }
    }

    return Array.from(result);
};

console.log("RESULT " + intersection([1,2,2,1], [2,2]))
