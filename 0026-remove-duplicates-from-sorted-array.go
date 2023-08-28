package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {
	p1, p2, l := 0, 1, len(nums)
	// p1 定位交换后数组内最大的数字，p2负责寻找p1后第一个大于
	for p1 < l && p2 < l {
		if nums[p1] != nums[p2] {
			nums[p1+1] = nums[p2]
			p1++
			p2++
		} else {
			// 找到p1后第一个和p1不等的index
			p2++
		}

	}

	return p1 + 1

}
