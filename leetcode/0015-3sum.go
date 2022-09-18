package leetcode

import "sort"

// 双指针
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			return res //剪枝
		}
		if i > 0 && n1 == nums[i-1] {
			continue //去重
		}
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for l < r && nums[l+1] == n2 {
					l++ //去重
				}
				for l < r && nums[r-1] == n3 {
					r-- //去重
				}
				l++
				r--
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
