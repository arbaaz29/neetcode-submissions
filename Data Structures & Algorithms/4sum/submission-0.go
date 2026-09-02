func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    res := [][]int{}
    quad := []int{}

    var kSum func(k, start, target int)
    kSum = func(k, start, target int) {
        if k == 2 {
            l, r := start, len(nums)-1
            for l < r {
                sum := nums[l] + nums[r]
                if sum < target {
                    l++
                } else if sum > target {
                    r--
                } else {
                    temp := make([]int, len(quad))
                    copy(temp, quad)
                    temp = append(temp, nums[l], nums[r])
                    res = append(res, temp)
                    l++
                    r--
                    for l < r && nums[l] == nums[l-1] {
                        l++
                    }
                    for l < r && nums[r] == nums[r+1] {
                        r--
                    }
                }
            }
            return
        }

        for i := start; i <= len(nums)-k; i++ {
            if i > start && nums[i] == nums[i-1] {
                continue
            }
            quad = append(quad, nums[i])
            kSum(k-1, i+1, target-nums[i])
            quad = quad[:len(quad)-1]
        }
    }

    kSum(4, 0, target)
    return res
}