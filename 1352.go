type ProductOfNumbers struct {
	p []int
}

func Constructor() ProductOfNumbers {
	var o ProductOfNumbers
	o.p = make([]int, 0)
	o.p = append(o.p, 1)
	return o
}

func (o *ProductOfNumbers) Add(num int) {
	if num == 0 {
		o.p = o.p[:0]
		o.p = append(o.p, 1)
	} else {
		var n, v int
		n = len(o.p)
		v = o.p[n-1] * num
		o.p = append(o.p, v)
	}
}

func (o *ProductOfNumbers) GetProduct(k int) int {
	if k >= len(o.p) {
		return 0
	}
	var n int
	n = len(o.p)
	return o.p[n-1] / o.p[n-1-k]
}
