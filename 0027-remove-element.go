package main

// 27. 移除元素
// https://leetcode.cn/problems/remove-element/
// 学习了如何使用指针？第一从用捏

func main() {
	removeElement([]int{3, 2, 2, 3}, 3)
}

func removeElement(nums []int, val int) int {
	l := len(nums)
	p1, p2 := 0, 0
	// p1指向首个=val的元素， p2指向p1后非val的第一个元素

Loop:
	for p1 < l && p2 < l {
		// 如果p1不是val， p1后移
		if nums[p1] != val {
			p1++
		} else {
			// 如果p1指向了val，p2先和p2同步，再后移找到第一个非val index
			p2 = p1
			for nums[p2] == val {
				p2++
				if p2 == l {
					goto Loop
				}
			}

			swap(&nums, p1, p2)
		}
	}
	return p1
}

func swap(nums *[]int, index1 int, index2 int) {
	temp := (*nums)[index1]
	(*nums)[index1] = (*nums)[index2]
	(*nums)[index2] = temp
}
