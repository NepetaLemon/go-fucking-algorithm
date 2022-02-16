package two_sum_1

func twoSum(nums []int, target int) []int {
	memoDiff := make(map[int]int)
	// 把索引和值倒过来，方便判断所需差值是否存在，以及它对应的索引
	for k, v := range nums {
		memoDiff[v] = k
	}
	for k, v := range nums {
		// 题目确定只有唯一解
		// 可以直接通过 hash 直接索引到所需差值是否存在
		idx, ok := memoDiff[target-v]
		// 一定要不是同一个索引值
		if ok && k != idx {
			return []int{k, idx}
		}
	}

	return []int{}
}
