func minSumOfLengths(arr []int, target int) int {
    INF := int(1e9)

    get_sub_arrays := func(arr []int) []int {
        lookup := map[int]int{}
        running_sum := 0
        dp := make([]int, len(arr))
        for i := range dp {
            dp[i] = INF
        }

        for i, num := range arr {
            running_sum += num
            if running_sum == target {
                dp[i] = i - 0 + 1
            } else if prev, ok := lookup[running_sum-target]; ok {
                dp[i] = i - prev + 1
            }
            lookup[running_sum] = i + 1

            if i > 0 && dp[i-1] < dp[i] {
                dp[i] = dp[i-1]
            }
        }
        return dp
    }

    dp_left := get_sub_arrays(arr)

    reversed_arr := make([]int, len(arr))
    for i := range arr {
        reversed_arr[i] = arr[len(arr)-1-i]
    }

    dp_right_rev := get_sub_arrays(reversed_arr)
    dp_right := make([]int, len(arr))
    for i := range dp_right_rev {
        dp_right[i] = dp_right_rev[len(dp_right_rev)-1-i]
    }

    ans := INF
    for i := 1; i < len(arr); i++ {
        if dp_left[i-1]+dp_right[i] < ans {
            ans = dp_left[i-1] + dp_right[i]
        }
    }

    if ans != INF {
        return ans
    }
    return -1
}
