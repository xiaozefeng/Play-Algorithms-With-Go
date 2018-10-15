package basic

// bubble sort
func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ { // 外层循环表示需要遍历多少趟
		for j := 0; j < n-1-i; j++ { // 内层循环表示每趟需要比较的次数
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 优化：鸡尾酒排序
func CocktailSort(arr []int) []int {
	n := len(arr)
	j, left, right := 0, 0, n-1
	for left < right {
		for j = left; j < right; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		right--

		for j = right; j > left; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
		left++
	}
	return arr
}
