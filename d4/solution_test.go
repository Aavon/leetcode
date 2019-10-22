package d4

import (
	"fmt"
	"testing"
)

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个有序数组的中位数
 *
 * https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (36.11%)
 * Likes:    1628
 * Dislikes: 0
 * Total Accepted:    101.5K
 * Total Submissions: 281K
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
 *
 * 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
 *
 * 你可以假设 nums1 和 nums2 不会同时为空。
 *
 * 示例 1:
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * 则中位数是 2.0
 *
 *
 * 示例 2:
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * 则中位数是 (2 + 3)/2 = 2.5
 *
 *
 */

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	c1 := len(nums1)
	c2 := len(nums2)
	mid := (c1+c2) / 2 + 1
	midVals := [2]int{}
	var idx0, idx1, idx2 int
	var val int
	for ;idx0 < mid; {
		if idx1 < c1 && (idx2 >= c2 || nums1[idx1] < nums2[idx2]) {
			val = nums1[idx1]
			idx1++
		} else {
			val = nums2[idx2]
			idx2++
		}
		midVals[idx0%2] = val
		idx0++
	}
	if (c1 + c2) % 2 == 0 {
		return (float64(midVals[0]) + float64(midVals[1])) / 2
	} else {
		if mid % 2 == 0 {
			return float64(midVals[1])
		} else {
			return float64(midVals[0])
		}
	}
}

func Test_01(t *testing.T) {
	nums1 := []int{1,3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}