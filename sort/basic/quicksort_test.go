package basic

import (
	"github.com/xiaozefeng/Play-Algorithms-With-Go/sort/arrayutil"
	"testing"
)

func TestQuickSort(t *testing.T) {
	s:= arrayutil.GenerateNearlyOrderArray(100000, 10)
	//s := []int{3, 2, 1, 0}
	arrayutil.InvokeTestNoReturnValue("quickSort", s, QuickSort)
}
