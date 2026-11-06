func findKthSmallest(coins []int, k int) int64 {
    sort.Ints(coins)

    gcd := func(a, b int64) int64 {
        for b != 0 {
            a, b = b, a%b
        }
        return a
    }

    count := func(m int64) int64 {
        n := len(coins)
        var total int64

        for mask := 1; mask < (1 << n); mask++ {
            lcm := int64(1)
            bits := 0

            for i, c := range coins {
                if mask>>i&1 == 0 {
                    continue
                }

                bits++
                cv := int64(c)
                g := gcd(lcm, cv)

                if lcm/g > m/cv {
                    lcm = m + 1
                    break
                }

                lcm = lcm / g * cv
            }

            if lcm <= m {
                if bits%2 == 1 {
                    total += m / lcm
                } else {
                    total -= m / lcm
                }
            }
        }

        return total
    }

    left, right := int64(1), int64(coins[0])*int64(k)

    for left < right {
        mid := left + (right-left)/2
        if count(mid) >= int64(k) {
            right = mid
        } else {
            left = mid + 1
        }
    }

    return left
}
