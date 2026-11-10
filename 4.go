import "math"

func findMedianSortedArrays(a []int, b []int) float64 {
    if len(a) > len(b) {
        a, b = b, a
    }
    m, n := len(a), len(b)
    tot := (m + n + 1) / 2
    l, r := 0, m
    for l <= r {
        i := (l + r) / 2
        j := tot - i
        l1, r1 := math.MinInt64, math.MaxInt64
        if i > 0 {
            l1 = a[i-1]
        }
        if i < m {
            r1 = a[i]
        }
        l2, r2 := math.MinInt64, math.MaxInt64
        if j > 0 {
            l2 = b[j-1]
        }
        if j < n {
            r2 = b[j]
        }
        if l1 <= r2 && l2 <= r1 {
            if (m+n)%2 == 1 {
                if l1 > l2 {
                    return float64(l1)
                }
                return float64(l2)
            }
            lm := l1
            if l2 > lm {
                lm = l2
            }
            rm := r1
            if r2 < rm {
                rm = r2
            }
            return (float64(lm) + float64(rm)) / 2.0
        } else if l1 > r2 {
            r = i - 1
        } else {
            l = i + 1
        }
    }
    return 0.0
}
