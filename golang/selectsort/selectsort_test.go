package selectsort

import "testing"

func TestSelectSort(t *testing.T) {
	var (
		in = []int{10, 2, 7, 4, 1, 2}
		expect = []int{1, 2, 2, 4, 7, 10}
	)

	actual := SelectSort(in)
	if ! IsSliceEqual(actual, expect) {
		t.Errorf("SelectSort(%d) is %d, but expect is %d", in, actual, expect)
	}
}

func IsSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil ) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

func BenchmarkQuickSortByInPlace(b *testing.B) {
	var in = []int{10, 2, 7, 4, 1, 2}
	for n := 0; n < b.N; n++ {
		SelectSort(in)
	}
}