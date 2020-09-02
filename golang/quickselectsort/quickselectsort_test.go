package quickselectsort

import "testing"

func TestQuickSelectSortByDC(t *testing.T) {
	var (
		inputArray = []int{10, 2, 7, 4, 1, 2, 20}
		inputK  = 5
		expectOutput  = 10
	)

	actual := QuickSelectSortByDC(inputArray, inputK)
	if  actual[len(actual)-1] != expectOutput {
		t.Errorf("QuickSelectSortByDC(%d, %d) top-K is %d, but expect is %d", inputArray, inputK, actual[len(actual)-1], expectOutput)
	}
}

func BenchmarkQuickSelectSortByDC(b *testing.B) {
	var (
		inputArray = []int{10, 2, 7, 4, 1, 2, 20}
		inputK  = 5
	)
	for n := 0; n < b.N; n++ {
		QuickSelectSortByDC(inputArray, inputK)
	}
}