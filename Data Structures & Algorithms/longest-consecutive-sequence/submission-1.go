func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
        return 0
    }
	// sort so that consecutive sequence is identified easily
    sort.Ints(nums)
	//Contains the latest streak after iterating the array
    res := 0
	//while the loop is running
	//curr will hold the current value, streak will hold current streak
    curr, streak := nums[0], 0
	// loop over the elements of the array
	i := 0
    for  i < len(nums) {
		// if curr element is not equal to the next element reset the streak and replace current number with the `i`th value of nums
		//this is the base condition that will reset and maintain streak 
        if curr != nums[i] {
            curr = nums[i]
            streak = 0
        }
		// Increment i if the `i`th nums value and curr value is same (duplicates) and come out of loop when the condition is false
        for i < len(nums) && nums[i] == curr {
            i++
        }
		// i will contain the latest index after ignoring all the duplicates
		// Increment the streak
        streak++
		// Increment the current value, now `curr` will represent the next expected consecutive number
        curr++
		// If current streak is greater than latest streak, update the latest streak
        if streak > res {
            res = streak
        }
    }
	// return the latest streak
    return res
}
