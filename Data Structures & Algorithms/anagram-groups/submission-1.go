import "slices"
func groupAnagrams(strs []string) [][]string {
    hm := make(map[string][]string)
    res := [][]string{}
 for _,val:= range strs{
    sortedS := sortString(val)
    hm[sortedS] = append(hm[sortedS],val)
 }

 for _, st:= range hm{
    res = append(res,st)
 }
 return res
}

func sortString(st string) string{
    str := []rune(st)
    slices.Sort(str)
    return string(str)
}