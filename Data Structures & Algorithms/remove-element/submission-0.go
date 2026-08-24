func removeElement(nums []int, val int) int {
    j := 0
    for idx:=0; idx<len(nums);idx++{
        if val != nums[idx]{
            nums[j] = nums[idx]
            j++
        }
    }
    return j
}
