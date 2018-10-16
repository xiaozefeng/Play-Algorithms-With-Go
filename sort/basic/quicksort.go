package basic

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}

// 对arr[l...r]前闭后闭快速排序
func quickSort(values []int, l, r int) {
	if l >= r {
		return
	}

	p := partition(values, l, r)
	quickSort(values, l, p-1)
	quickSort(values, p+1, r)
}

// 对arr[l...r]部分进行partition操作
// 返回p 使得arr[l ...p-1] < arr[p] ; arr[p+1...r] >arr[p]
func partition(values []int, l, r int) int {
	v := values[l]
	j := l
	for i := l + 1; i <= r; i++ {
		if values[i] < v {
			values[j+1], values[i] = values[i], values[j+1]
			j++
		}
	}
	values[l], values[j] = values[j], values[l]
	return j
}

// 第二种实现
func QSort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	left, right := 0, len(data)-1
	for i := 1; i <= right; {
		if data[i] > mid {
			data[i], data[right] = data[right], data[i]
			right--
		} else {
			data[i], data[left] = data[left], data[i]
			left++
			i++
		}
	}
	QSort(data[:left])
	QSort(data[left+1:])
}
