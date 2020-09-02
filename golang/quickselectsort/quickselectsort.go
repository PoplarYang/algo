package quickselectsort

// resolve Top-k problem
func QuickSelectSortByDC(lx []int, k int) (ls []int) {
	if len(lx) < 2 {
		return lx
	}
	pivot := lx[0]
	left := make([]int, 0)
	right := make([]int, 0)
	for _, v := range lx[1:] {
		if pivot < v {
			right = append(right, v)
		} else {
			left = append(left, v)
		}
	}
	leftLength := len(left)
	if leftLength == k {
		return append(left, pivot)
	} else if leftLength < k {
		return QuickSelectSortByDC(left, k)
	} else {
		return QuickSelectSortByDC(right, k-leftLength)
	}
}