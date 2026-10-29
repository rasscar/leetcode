func findItinerary(tickets [][]string) []string {
	var i, j int
	var tujuan map[string][]string
	tujuan = make(map[string][]string)

	for i = 0; i < len(tickets); i++ {
		tujuan[tickets[i][0]] = append(tujuan[tickets[i][0]], tickets[i][1])
	}

	for k := range tujuan {
		var arr []string
		arr = tujuan[k]
		for i = 1; i < len(arr); i++ {
			var key string
			key = arr[i]
			j = i - 1
			for j >= 0 && arr[j] < key {
				arr[j+1] = arr[j]
				j--
			}
			arr[j+1] = key
		}
	}

	var tumpukan []string
	tumpukan = make([]string, 0)
	tumpukan = append(tumpukan, "JFK")

	var rute []string
	rute = make([]string, 0)

	for len(tumpukan) > 0 {
		var skr string
		skr = tumpukan[len(tumpukan)-1]

		if len(tujuan[skr]) > 0 {
			var next string
			next = tujuan[skr][len(tujuan[skr])-1]
			tujuan[skr] = tujuan[skr][:len(tujuan[skr])-1]
			tumpukan = append(tumpukan, next)
		} else {
			rute = append(rute, skr)
			tumpukan = tumpukan[:len(tumpukan)-1]
		}
	}

	var hasil []string
	hasil = make([]string, len(rute))
	for i = 0; i < len(rute); i++ {
		hasil[i] = rute[len(rute)-1-i]
	}

	return hasil
}
