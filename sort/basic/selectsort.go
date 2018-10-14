package basic

func SelectionSort(arr []int) []int {
	n:= len(arr)
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = swap(arr[i], arr[minIndex])
	}
	return arr
}

func swap(i, j int) (int, int) {
	return j, i
}
