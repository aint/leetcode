function pivotIndex(nums: number[]): number {
    for (let i = 0; i < nums.length; i++) {

        let leftSum = 0;
        for (let j = 0; j <= i; j++) {
            leftSum += nums[j];
        }

        let rightSum = 0;
        for (let j = nums.length - 1; j >= i; j--) {
            rightSum += nums[j];
        }

        if (leftSum == rightSum) {
            return i;
        }
    }
    return -1
};

console.log("RESULT " + pivotIndex([1,7,3,6,5,6]))
