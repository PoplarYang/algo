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
	length := len(lx)
	if length < 2 {
		return lx
	}

	//for k1, _ := range lx {
	//	for k2, _ := range lx[k1:] {
	//		if lx[k1] > lx[k1+k2] {
	//			lx[k1], lx[k1+k2] = lx[k1+k2], lx[k1]
	//		}
	//	}
	//}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if lx[i] > lx[j] {
				lx[i], lx[j] = lx[j], lx[i]
			}
		}
	}
	return lx
}

func QuickSortByThreeCenter(lx []int) (ls []int) {
	if len(lx) < 2 {
		return lx
	}
	pivotIndex := getPivot(lx)
	pivot := lx[pivotIndex]
	less := make([]int, 0)
	greater := make([]int, 0)
	lx = append(lx[:pivotIndex], lx[pivotIndex+1:]...)
	for _, v := range  lx {
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

func getPivot(lx []int) int {
	length := len(lx)
	if length < 3 {
		return lx[0]
	}
	middle := length / 2
	if lx[0] <= lx[middle]  && lx[middle] <= lx[length-1] || lx[length-1] <= lx[middle]  && lx[middle] <= lx[0] {
		return middle
	} else if lx[middle] <= lx[0]  && lx[0] <= lx[length-1] || lx[length-1] <= lx[0]  && lx[0] <= lx[middle] {
		return 0
	//} else lx[0] <= lx[length-1]  && lx[length-1] <= lx[middle] || lx[middle] <= lx[length-1]  && lx[length-1] <= lx[0] {
	} else {
		return length-1
	}
}