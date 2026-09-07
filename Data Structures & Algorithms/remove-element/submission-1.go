func removeElement(nums []int, val int) int {
    el := make([]int, len(nums))
    j := 0
    for i:=0;i<len(nums);i++{
        if nums[i]!=val{
            el[j] = nums[i]
            j++
        }
    }
    for i:=0;i<len(nums);i++{
        nums[i]=el[i]
    }
    return j
    
}
