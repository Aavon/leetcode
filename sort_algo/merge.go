package sort_algo

// 归并排序
// 稳定
// nlogn

func MergeSort(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	}
	left := MergeSort(arr[:l/2])
	right := MergeSort(arr[l/2:])
	mergedLen := len(left) + len(right)
	merged := make([]int, mergedLen)
	li := 0
	ri := 0
	for mi := 0; mi < mergedLen; mi++ {
		if li == len(left) || (ri != len(right) && right[ri] < left[li]) {
			merged[mi] = right[ri]
			ri++
		} else {
			merged[mi] = left[li]
			li++
		}

	}
	return merged
}
