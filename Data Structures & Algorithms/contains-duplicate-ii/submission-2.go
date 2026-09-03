func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]bool)
	l := 0
	for r:=0;r<len(nums);r++{
		if r-l >k{
			delete(seen, nums[l])
			l++
		}
		if seen[nums[r]]{
			return true
		}
		seen[nums[r]] = true
	} 
	return false
}
