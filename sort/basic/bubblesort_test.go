package basic

import (
	"github.com/xiaozefeng/Play-Algorithms-With-Go/sort/arrayutil"
	"testing"
)

func TestBubbleSort2(t *testing.T) {
	s := []int{3, 2, 1, 0}
	BubbleSort2(s)
}

func TestBubbleSort(t *testing.T) {
	s := []int{3, 2, 1, 0}
	BubbleSort(s)
}

func TestPerformance(t *testing.T) {
	n := 20000
	s1 := arrayutil.GenerateRandomArray(n, n)
	s2 := make([]int, n)
	copy(s2, s1)
	s3 := make([]int, n)
	copy(s3, s1)
	arrayutil.InvokeTestNoReturnValue("BubbleSort1", s1, BubbleSort)
	// 最快的是2
	arrayutil.InvokeTestNoReturnValue("BubbleSort2", s2, BubbleSort2)
	arrayutil.InvokeTestNoReturnValue("BubbleSort3", s3, BubbleSort3)
}
