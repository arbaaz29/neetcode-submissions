func majorityElement(nums []int) int {
    freq := make(map[int]int)
    for _,val:=range nums{
        freq[val]++
    }
    champ := 0
    for idx,val := range freq{
        if val > len(nums)/2{
            champ = idx
        }
    }
    return champ
}
