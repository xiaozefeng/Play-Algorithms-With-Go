package arrayutil

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成随机数组
func GenerateRandomArray(length, rangeR int) []int {
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = rand.Intn(rangeR)
	}
	return result
}

// 生成一个近乎有序的数组
func GenerateNearlyOrderArray(length, swapTime int) []int {
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = i
	}
	for i := 0; i < swapTime; i++ {
		a := rand.Intn(length)
		b := rand.Intn(length)
		result[a], result[b] = result[b], result[a]
	}
	return result
}

// 判断数组是否有序
func IsSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

// 执行测试
func InvokeTest(fName string, arr []int, f func(arr []int) []int) {
	now := time.Now()
	arr = f(arr)
	since := time.Since(now)
	if IsSorted(arr) {
		fmt.Printf("%s执行时间:%v 秒\n", fName, since.Seconds())
	} else {
		fmt.Println("数组无序")
	}
}

// 执行测试
func InvokeTestNoReturnValue(fName string, arr []int, f func(arr []int)) {
	now := time.Now()
	f(arr)
	since := time.Since(now)
	if IsSorted(arr) {
		fmt.Printf("%s执行时间:%v 秒\n", fName, since.Seconds())
	} else {
		fmt.Println("数组无序")
	}
}
