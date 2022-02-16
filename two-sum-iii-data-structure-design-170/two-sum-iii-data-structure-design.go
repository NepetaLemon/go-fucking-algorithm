package two_sum_iii_data_structure_design_170

type TwoSum struct {
	items []int
	m     map[int]int
}

func Constructor() TwoSum {
	return TwoSum{
		items: make([]int, 0),
		m:     make(map[int]int),
	}
}

// 这题不像 1 ，是可以明确所求答案的，如果要提升 find 的效率，只能先暴力算出所有可能结果，牺牲空间，节省时间（但是这题超时）
func (this *TwoSum) Add(number int) {
	l := len(this.items)
	this.items = append(this.items, number)
	this.m[number] = l
}

func (this *TwoSum) Find(value int) bool {
	for k, v := range this.items {
		if idx, ok := this.m[value-v]; ok && k != idx {
			return true
		}
	}

	return false
}
