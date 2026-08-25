func subsetXORSum(nums []int) int {
    totalSum := 0

    var backtrack func(index int, currentXOR int)
    backtrack = func(index int, currentXOR int) {
        if index == len(nums) {
            totalSum += currentXOR
            return
        }

        backtrack(index+1, currentXOR^nums[index])

        backtrack(index+1, currentXOR)
    }

    backtrack(0, 0)
    return totalSum
}