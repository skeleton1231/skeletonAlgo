package leetcode

func search(nums []int, target int) int {

	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)>>1
		for left < right {
			if nums[mid] > target {
				right = mid - 1
			} else if nums[mid] < target {
				left = mid + 1
			} else {
				return mid
			}
		}
	}
	return -1
}
