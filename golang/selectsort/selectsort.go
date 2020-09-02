package selectsort

func SelectSort(lx []int) (ls []int){
	length := len(lx)
	if length <= 1 {
		return lx
	}

	for i := 0; i < length; i++ {
		// 初始的最小值位置从0开始，依次向右
		min_index := i

		// 从i右侧的所有元素中找出当前最小值所在的下标
		for j := i + 1; j < length; j++ {
			if lx[j] < lx[min_index] {
				min_index = j
			}
		}
		// 如果选择出的数据不在左边的位置，进行交换
		if min_index != i {
			// 把每次找出来的最小值与之前的最小值做交换
			lx[i], lx[min_index] = lx[min_index], lx[i]
		}
	}
	return lx
}