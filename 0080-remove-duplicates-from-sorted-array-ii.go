package main

// 猪脑过载，太难了
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/

func removeDuplicatesii(nums []int) int {

	removek := func(k int) int {
		i := 0                   // 定义指针 i，写数组的指针
		for _, v := range nums { // 遍历数组

			// 在第 i 个元素的值 等于 第 i - k 个元素的时候，说明当前值超过 k 个了，所以要覆盖掉第 i 个元素
			// 一直往后遍历数组直到出现第一个值使得 nums[i] != nums[i - k]

			if i < k || v != nums[i-k] {
				// 只有在出现使得元素 nums[i] 出现小于 k 次的元素时，写指针才会后移！
				nums[i] = v
				i++

			}
		}

		return i

	}

	return removek(2)
}
