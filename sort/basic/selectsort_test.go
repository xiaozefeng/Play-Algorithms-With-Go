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
	s3 := make([]int, n)
	copy(s3, s1)
	s4 := make([]int, n)
	copy(s4, s1)
	s5 := make([]int, n)
	copy(s5, s1)
	arrayutil.InvokeTest("SelectionSort", s2, SelectionSort)
	arrayutil.InvokeTest("InsertionSort", s1, InsertionSort)
	arrayutil.InvokeTestNoReturnValue("BubbleSort", s3, BubbleSort)
	arrayutil.InvokeTest("CocktailSort", s4, CocktailSort)
	arrayutil.InvokeTestNoReturnValue("QuickSort", s5, QuickSort)
}
