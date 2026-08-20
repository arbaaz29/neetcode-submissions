func hasDuplicate(nums []int) bool {
	for i, _ := range nums{
		for j:= i+1; j < len(nums); j++ {
			if nums[i] == nums[j]{
					return true
			}
		}
	}
	return false
}
