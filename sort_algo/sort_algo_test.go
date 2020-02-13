package sort_algo

import (
	"fmt"
	"math/rand"
	"testing"
)

// https://images2017.cnblogs.com/blog/849589/201710/849589-20171015233043168-1867817869.png

const length = 10

var data = make([]int, 0, length)

func init() {
	//rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		data = append(data, rand.Intn(1000))
	}
}

func TestBubbleSort(t *testing.T) {
	fmt.Println(BubbleSort(data))
}

func TestSelectionSort(t *testing.T) {
	fmt.Println(SelectionSort(data))
}

func TestInsertSort(t *testing.T) {
	fmt.Println(InsertSort(data))
}

func TestShellSort(t *testing.T) {
	fmt.Println(ShellSort(data))
}

func TestQuickSort(t *testing.T) {
	fmt.Println(QuickSort(data))
}

func TestMergeSort(t *testing.T) {
	fmt.Println(MergeSort(data))
}

func TestHeapSort(t *testing.T) {
	fmt.Println(HeapSort(data))
}
