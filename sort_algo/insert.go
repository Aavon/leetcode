package sort_algo

// 插入排序
// 稳定
// n2

func InsertSort(arr []int) []int {
	l := len(arr)
	for bottomN := 0; bottomN < l-1; bottomN++ {
		for next := bottomN + 1; next > 0; next-- {
			if arr[next-1] > arr[next] {
				arr[next-1], arr[next] = arr[next], arr[next-1]
			} else {
				break
			}
		}
	}
	return arr
}
