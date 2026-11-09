func smallestNumber(num string, t int64) string {
    p2, p3, p5, p7, ok := faktorisasi(t)
    if !ok {
        return "-1"
    }

    dim2, dim3, dim5, dim7 := p2+1, p3+1, p5+1, p7+1
    totalStatus := dim2 * dim3 * dim5 * dim7

    indeks := func(e2, e3, e5, e7 int) int {
        return (((e2*dim3 + e3)*dim5 + e5)*dim7 + e7)
    }

    uraikan := func(status int) (int, int, int, int) {
        e7 := status % dim7
        status /= dim7
        e5 := status % dim5
        status /= dim5
        e3 := status % dim3
        e2 := status / dim3
        return e2, e3, e5, e7
    }

    vektorDigit := [10][4]int{
        {0, 0, 0, 0},
        {0, 0, 0, 0},
        {1, 0, 0, 0},
        {0, 1, 0, 0},
        {2, 0, 0, 0},
        {0, 0, 1, 0},
        {1, 1, 0, 0},
        {0, 0, 0, 1},
        {3, 0, 0, 0},
        {0, 2, 0, 0},
    }

    transisi := make([][9]int, totalStatus)
    for status := 0; status < totalStatus; status++ {
        e2, e3, e5, e7 := uraikan(status)
        for digit := 1; digit <= 9; digit++ {
            ne2 := e2 + vektorDigit[digit][0]
            if ne2 > p2 {
                ne2 = p2
            }
            ne3 := e3 + vektorDigit[digit][1]
            if ne3 > p3 {
                ne3 = p3
            }
            ne5 := e5 + vektorDigit[digit][2]
            if ne5 > p5 {
                ne5 = p5
            }
            ne7 := e7 + vektorDigit[digit][3]
            if ne7 > p7 {
                ne7 = p7
            }
            transisi[status][digit-1] = indeks(ne2, ne3, ne5, ne7)
        }
    }

    target := indeks(p2, p3, p5, p7)

    jarak := make([]int, totalStatus)
    for i := 0; i < totalStatus; i++ {
        jarak[i] = -1
    }
    jarak[target] = 0

    var telusuri func(int) int
    telusuri = func(status int) int {
        if jarak[status] != -1 {
            return jarak[status]
        }
        terbaik := 1 << 30
        for digit := 2; digit <= 9; digit++ {
            statusBaru := transisi[status][digit-1]
            if statusBaru == status {
                continue
            }
            nilai := telusuri(statusBaru)
            if nilai+1 < terbaik {
                terbaik = nilai + 1
            }
        }
        jarak[status] = terbaik
        return terbaik
    }
    telusuri(0)

    n := len(num)

    statusPrefix := make([]int, n+1)
    validPrefix := make([]bool, n+1)
    validPrefix[0] = true

    for i := 0; i < n; i++ {
        if !validPrefix[i] {
            continue
        }
        if num[i] == '0' {
            continue
        }
        digit := int(num[i] - '0')
        statusPrefix[i+1] = transisi[statusPrefix[i]][digit-1]
        validPrefix[i+1] = true
    }

    if validPrefix[n] && statusPrefix[n] == target {
        return num
    }

    buatSuffixMin := func(sisa int, statusAwal int) string {
        hasil := make([]byte, sisa)
        status := statusAwal
        for pos := 0; pos < len(hasil); pos++ {
            for digit := 1; digit <= 9; digit++ {
                statusBaru := transisi[status][digit-1]
                if jarak[statusBaru] != -1 && jarak[statusBaru] <= sisa-1 {
                    hasil[pos] = byte('0' + digit)
                    status = statusBaru
                    sisa--
                    break
                }
            }
        }
        return string(hasil)
    }

    for i := n - 1; i >= 0; i-- {
        if !validPrefix[i] {
            continue
        }
        status := statusPrefix[i]
        sisa := n - i - 1
        rendah := int(num[i]-'0') + 1
        if rendah < 1 {
            rendah = 1
        }
        for digit := rendah; digit <= 9; digit++ {
            statusBaru := transisi[status][digit-1]
            if jarak[statusBaru] != -1 && jarak[statusBaru] <= sisa {
                suffix := buatSuffixMin(sisa, statusBaru)
                return num[:i] + string(byte('0'+digit)) + suffix
            }
        }
    }

    panjang := n + 1
    if jarak[0] > panjang {
        panjang = jarak[0]
    }
    return buatSuffixMin(panjang, 0)
}

func faktorisasi(x int64) (int, int, int, int, bool) {
    p2 := 0
    for x%2 == 0 {
        p2++
        x /= 2
    }
    p3 := 0
    for x%3 == 0 {
        p3++
        x /= 3
    }
    p5 := 0
    for x%5 == 0 {
        p5++
        x /= 5
    }
    p7 := 0
    for x%7 == 0 {
        p7++
        x /= 7
    }
    return p2, p3, p5, p7, x == 1
}
