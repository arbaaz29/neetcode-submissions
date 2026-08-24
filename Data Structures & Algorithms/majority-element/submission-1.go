func majorityElement(nums []int) int {
    sort.Ints(nums)
    n:=len(nums)
    for _, v := range nums{
        count := 0
        for _, i := range nums{
            if i==v{
                count++
            }
        }
        if count > n/2 {
            return v
        }
    }
    return -1
}
