func majorityElement(nums []int) int {
    freq := make(map[int]int)
    champ := 0
    for _,val:=range nums{
        freq[val]++
        if freq[val] > len(nums)/2{
            champ = val
        }
    }
    return champ
}
