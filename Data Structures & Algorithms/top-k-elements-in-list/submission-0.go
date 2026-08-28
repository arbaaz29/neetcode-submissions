func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	//create a slice of slice
	//eg. -> reps[0] → []int
	reps := make([][]int, len(nums)+1)
	// create a hash map of keys representing the number and values representing the frequency
	for _,val:= range nums{
		freq[val]++
	}
	// Create a array (bucket), the key (index) will the number of times a number has occured throughout the array, and the value (array) attached to the number will be the respective number.
	//eg. -> if 5 occured 3 times and 1 occured 3 times, the slice will look like following:
	// reps[3]=[1,5]
	for num,count:= range freq{
		reps[count] = append(reps[count], num)
	}
	//create empty slice, we will store the comparisions with k in res
	res := []int{}
	// iterate over all the rows in reps
	for i:=len(reps)-1;i>0;i--{
		// unpack the row and get the values from the respective columns
		for _,num := range reps[i]{
			res = append(res,num)
			// K defines the count of unique numbers to be returned
			//eg. - if k=2 that means we need to return the most frequent number and the 2nd most frequent number
			if len(res)==k{
				return res
			}
		}
	}
	return res
}
