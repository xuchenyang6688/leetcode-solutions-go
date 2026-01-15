package backtrack

func permutate(nums []int) [][]int {
	result := [][]int{}
	backtrack(nums, 0, &result)
	return result
}

func backtrack(nums []int, index int, result *[][]int) {
    if index == len(nums) {
		permutation :=make([]int, len(nums))
		copy(permutation, nums)
		*result = append(*result, permutation)
		return
	}
	
	for i := index; i<len(nums); i++ {
		nums[i], nums[index] = nums[index], nums[i]
		backtrack(nums, index+1, result)
		nums[i], nums[index] = nums[index], nums[i]
	}
}