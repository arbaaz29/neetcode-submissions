func findClosestElements(arr []int, k int, x int) []int {
    n := len(arr)
    idx := 0
    for i := 1; i < n; i++ {
        if abs(x-arr[idx]) > abs(x-arr[i]) {
            idx = i
        }
    }

    res := []int{arr[idx]}
    l, r := idx-1, idx+1

    for len(res) < k {
        if l >= 0 && r < n {
            if abs(x-arr[l]) <= abs(x-arr[r]) {
                res = append(res, arr[l])
                l--
            } else {
                res = append(res, arr[r])
                r++
            }
        } else if l >= 0 {
            res = append(res, arr[l])
            l--
        } else if r < n {
            res = append(res, arr[r])
            r++
        }
    }

    sort.Ints(res)
    return res
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}