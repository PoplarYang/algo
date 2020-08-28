package quicksort

func QuickSortByDC(lx []int) (ls []int) {
	if len(lx) < 2 {
		return lx
	}
	pivot := lx[0]
	less := make([]int, 0)
	greater := make([]int, 0)
	for _, v := range lx[1:] {
		if pivot < v {
			greater = append(greater, v)
		} else {
			less = append(less, v)
		}
	}
	ls = append(QuickSortByDC(less), pivot)
	ls = append(ls, QuickSortByDC(greater)...)
	return ls
}

func QuickSortByInPlace(lx []int) []int {
	if len(lx) < 2 {
		return lx
	}

	for k1, _ := range lx {
		for k2, _ := range lx[k1:] {
			if lx[k1] > lx[k1+k2] {
				lx[k1], lx[k1+k2] = lx[k1+k2], lx[k1]
				//break
			}
		}
	}
	return lx
}