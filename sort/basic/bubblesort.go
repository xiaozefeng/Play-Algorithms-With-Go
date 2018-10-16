package basic

// bubble sort
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ { // 外层循环表示需要遍历多少趟
		for j := 0; j < n-1-i; j++ { // 内层循环表示每趟需要比较的次数
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 优化算法2
func BubbleSort2(arr []int) {
	n := len(arr)
	swapped := true
	for swapped {
		swapped = false
		// 优化, 每一趟Bubble Sort都将最大的元素放在了最后的位置
		// 所以下一次排序, 最后的元素可以不再考虑
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
		n--
	}
}

// 优化算法3
func BubbleSort3(arr []int) {
	n := len(arr)
	newN := 1
	for newN > 0 {
		newN = 0
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				// 记录最后一次交互的位置, 在此之后的元素在下一轮的扫描中不予考虑
				newN = i
			}
		}
		n = newN
	}
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
