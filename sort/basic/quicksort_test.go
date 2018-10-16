package basic

import (
	"github.com/xiaozefeng/Play-Algorithms-With-Go/sort/arrayutil"
	"testing"
)

func TestQuickSort(t *testing.T) {
	s:= arrayutil.GenerateRandomArray(200000, 1000000)
	//s := []int{3, 2, 1}
	arrayutil.InvokeTestNoReturnValue("quickSort", s, QuickSort)
}

func TestQSort(t *testing.T) {
	s := arrayutil.GenerateRandomArray(2000000, 1000000)
	arrayutil.InvokeTestNoReturnValue("QuickSort2", s, QSort)
}