package leetcode

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)
	var ants []int
	for i, v := range nums {
		res := target - v
		if pos, ok := m[res]; ok {
			ants = []int{pos, i}
			break
		}
		m[v] = i
	}
	return ants
}
