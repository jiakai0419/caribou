package sliceutil

func SplitN(s []int64, n int) [][]int64 {
	l := len(s)/n
	if len(s)%n != 0 {
		l++
	}
	var ss [][]int64 = make([][]int64, l, l)
	row := 0
	count := 0
	for _, v := range s {
		if count!= 0 && count % n == 0 {
			row++
			count = 0
		}
		ss[row] = append(ss[row], v)
		count++
	}
	return ss
}
