func sortArray(nums []int) []int {
    n := len(nums)
	if n ==1{
		return nums
	}
	shellSort(nums, n)
	return nums
}

func shellSort(nums []int, n int){
	gap := n/2
	for gap >= 1{
		for  i:= gap; i < n; i++{
			tmp := nums[i]
			j := i-gap
			for j>=0 && nums[j] > tmp{
				nums[j+gap] = nums[j]
				j -=gap
			}
			nums[j+gap] = tmp
		}
		gap /=2
	}
}
