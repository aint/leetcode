function pivotIndex(nums: number[]): number {
    let leftSum = 0;
    let rightSum = 0;
    for (let i = 0; i < nums.length; i++) {
        rightSum += nums[i];
    }

    for (let i = 0; i < nums.length; i++) {
        leftSum += nums[i];

        if (leftSum == rightSum) {
            return i;
        }

        rightSum -= nums[i];
    }
    return -1
};

console.log("RESULT " + pivotIndex([1,7,3,6,5,6]))
