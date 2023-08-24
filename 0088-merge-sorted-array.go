package main

// 88. 合并两个有序数组
// https://leetcode.cn/problems/merge-sorted-array/

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 如此tricky的解法，这很难评！

	for p1, p2, p3 := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; p3-- {
		var cur int // 记录这一轮要被写入数组的数字

		if p1 < 0 {
			// nums1已经读取完了 直接读nums2
			cur = nums2[p2]
			p2--
		} else if p2 < 0 {
			cur = nums1[p1]
			p1--
		} else {
			//这种情况下都没有读完
			if nums1[p1] >= nums2[p2] {
				cur = nums1[p1]
				p1--
			} else {
				cur = nums2[p2]
				p2--
			}
		}

		nums1[p3] = cur

	}

}
