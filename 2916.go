const MOD int64 = 1_000_000_007

type Fenwick struct {
    n      int
    b1, b2 []int64
}

func NewFenwick(n int) *Fenwick {
    return &Fenwick{
        n:  n,
        b1: make([]int64, n+1),
        b2: make([]int64, n+1),
    }
}

func (f *Fenwick) add(bit []int64, idx int, val int64) {
    for i := idx; i <= f.n; i += i & -i {
        bit[i] += val
    }
}

func (f *Fenwick) update(i int, val int64) {
    f.add(f.b1, i, val)
    f.add(f.b2, i, int64(i-1)*val)
}

func (f *Fenwick) rangeAdd(l, r int, val int64) {
    f.update(l, val)
    f.update(r+1, -val)
}

func (f *Fenwick) query(bit []int64, idx int) int64 {
    var sum int64

    for i := idx; i > 0; i -= i & -i {
        sum += bit[i]
    }

    return sum
}

func (f *Fenwick) prefixSum(i int) int64 {
    s1 := f.query(f.b1, i)
    s2 := f.query(f.b2, i)

    return int64(i)*s1 - s2
}

func (f *Fenwick) rangeSum(l, r int) int64 {
    return f.prefixSum(r) - f.prefixSum(l-1)
}

func sumCounts(nums []int) int {
    n := len(nums)
    fw := NewFenwick(n)

    last := make(map[int]int)
    var sq, ans int64

    for r, x := range nums {
        p, ok := last[x]
        if !ok {
            p = -1
        }

        l, right := p+2, r+1

        oldSum := fw.rangeSum(l, right)
        length := int64(right - l + 1)

        sq = (sq + 2*oldSum + length) % MOD

        fw.rangeAdd(l, right, 1)

        ans = (ans + sq) % MOD

        last[x] = r
    }

    return int(ans)
}
