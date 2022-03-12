package backtracks

import (
	"sort"
)

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make(map[int]bool, 0)

	var backtrack func()
	backtrack = func() {
		// 如果当期的剩余集合中的元素为空，那么就将当前path加入结果
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}

		for i := 0; i < len(nums); i++ {
			// 如果已经将该元素添加到集合，则跳过
			if used[i] {
				continue
			}

			path = append(path, nums[i])
			used[i] = true

			backtrack()

			path = path[0 : len(path)-1]
			used[i] = false
		}
	}

	backtrack()

	return res
}

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make(map[int]bool, 0)

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
		}

		for i := 0; i < len(nums); i++ {
			if used[i] || i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			path = append(path, nums[i])
			used[i] = true

			backtrack()

			path = path[:len(path)-1]
			used[i] = false
		}
	}

	sort.Ints(nums)
	backtrack()

	return res
}