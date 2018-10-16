package basic

func InsertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		// 写法1
		//for j := i; j > 0; j-- {
		//	if arr[j] < arr[j-1] {
		//		arr[j], arr[j-1] = arr[j-1], arr[j]
		//	} else {
		//		break
		//	}
		//}

		// 写法2
		//for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
		//	arr[j], arr[j-1] = arr[j-1], arr[j]
		//}

		// 写法3 将交换位置改为赋值操作
		e := arr[i]
		j := i
		for ; j > 0 && arr[j-1] > e; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
	return arr
}
