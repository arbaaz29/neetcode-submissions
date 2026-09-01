func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i:=0; i < len(nums);i++{
		a := nums[i]
        if a > 0 {
            break
        }
        if i > 0 && a == nums[i-1] {
            continue
        }

		j,k :=i+1,len(nums)-1
		for j<k {
			sum := nums[i]+nums[j]+nums[k]
			if sum < 0 {
				j++
			} else if sum > 0{
				k--
			} else {
				res = append(res, []int{nums[i],nums[j],nums[k]})
				j++
				k--
				for j<k && nums[j] == nums[j-1]{
					j++
				}
			}
		}
	}
	return res
}
