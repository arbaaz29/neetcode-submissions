import "slices"
func isAnagram(s string, t string) bool {
	str1 := []rune(s)
	str2 := []rune(t)
	slices.Sort(str1)
	slices.Sort(str2)
	res1 := string(str1)
	res2 := string(str2)
	if res1 == res2{
		return true
	}
	return false
}
