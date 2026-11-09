func minimumPushes(word string) int {
var freq [26]int
for _, ch := range word {
freq[ch-'a']++
}
counts := make([]int, 0, 26)
for _, f := range freq {
if f > 0 {
counts = append(counts, f)
}
}
for i := 0; i < len(counts)-1; i++ {
for j := i + 1; j < len(counts); j++ {
if counts[i] < counts[j] {
counts[i], counts[j] = counts[j], counts[i]
}
}
}
total := 0
press := 1
key := 0
for _, f := range counts {
total += f * press
key++
if key == 8 {
press++
key = 0
}
}
return total
}
