package basic

func BubbleSort(arr []int) []int {
	n := len(arr)
	isSwap := true
	for i := 0; i < n-1 && isSwap; i++ {
		isSwap = false
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isSwap = true
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
