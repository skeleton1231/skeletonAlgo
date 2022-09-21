package leetcode

func removeElement(nums []int, val int) int {

	low := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[low] = nums[fast]
			low++
		}
	}
	nums = nums[:low]
	return low
}
