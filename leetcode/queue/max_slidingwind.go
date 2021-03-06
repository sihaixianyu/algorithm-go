package queue

func maxSlidingWindow(nums []int, k int) []int {
	var res, q []int
	push := func(i int) {
		for len(q) > 0 && nums[q[len(q)-1]] <= nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			push(i)
		} else {
			push(i)
			for q[0] <= i-k {
				q = q[1:]
			}
			res = append(res, nums[q[0]])
		}
	}

	return res
}
