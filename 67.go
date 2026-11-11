func addBinary(a string, b string) string {
    var bawa byte = '0'
    var hasil []byte
    i := len(a) - 1
    j := len(b) - 1

    for i >= 0 || j >= 0 || bawa == '1' {
        total := 0
        if i >= 0 && a[i] == '1' {
            total++
        }
        if j >= 0 && b[j] == '1' {
            total++
        }
        if bawa == '1' {
            total++
        }

        if total%2 == 1 {
            hasil = append(hasil, '1')
        } else {
            hasil = append(hasil, '0')
        }

        if total >= 2 {
            bawa = '1'
        } else {
            bawa = '0'
        }

        i--
        j--
    }

    for k := 0; k < len(hasil)/2; k++ {
        hasil[k], hasil[len(hasil)-1-k] = hasil[len(hasil)-1-k], hasil[k]
    }

    return string(hasil)
}
