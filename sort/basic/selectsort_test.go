package basic

import (
	"github.com/xiaozefeng/Play-Algorithms-With-Go/sort/arrayutil"
	"testing"
)

func TestPerformanceCompare(t *testing.T) {
	n := 20000
	s1 := arrayutil.GenerateRandomArray(n, 1000000)
	s2 := make([]int, n)
	copy(s2, s1)
	arrayutil.InvokeTest("SelectionSort", s1, SelectionSort)
	arrayutil.InvokeTest("InsertionSort", s2, InsertionSort)
}

func TestPerformanceCompare2(t *testing.T) {
	n := 20000
	s1 := arrayutil.GenerateNearlyOrderArray(n, 20)
	s2 := make([]int, n)
	copy(s2, s1)
	arrayutil.InvokeTest("SelectionSort", s2, SelectionSort)
	arrayutil.InvokeTest("InsertionSort", s1, InsertionSort)

}
