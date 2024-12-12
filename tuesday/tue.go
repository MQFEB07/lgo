package tuesday

// 2461. Maximum Sum of Distinct Subarrays With Length K
func maximumSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	if k <= 0 || k > n {
		return 0
	}
	frequency := make(map[int]int)
	start, sum, max_sum := 0, 0, 0
	//for loop till end reaches n
	for end := 0; end < n; end++ {
		frequency[nums[end]]++
		sum += nums[end]
		if end-start+1 == k {
			//distinct elements condition
			if len(frequency) == k {
				max_sum = max(max_sum, sum)
			}
			if frequency[nums[start]] == 1 {
				delete(frequency, nums[start])
			} else {
				frequency[nums[start]]--
			}
			sum -= nums[start]
			start++
		}
	}
	return int64(max_sum)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
