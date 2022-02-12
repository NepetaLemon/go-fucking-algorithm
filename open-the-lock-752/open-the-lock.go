package open_the_lock_752

// 出处：https://labuladong.gitee.io/algo/1/6/
// 题目：https://leetcode-cn.com/problems/open-the-lock/
// 这题有个特别的地方是「走回头路」，然后导致死循环，所以需要一个 map 来记录一下，已经测试过的组合
// 这里的 deadMap 可以融合到 visited 里面去，但是为了方便理解就不改了
func openLock(deadends []string, target string) int {
	initStr := "0000"
	visited := map[string]struct{}{initStr: {}}
	deadMap := getDeadEndsMap(deadends)
	queue := []string{initStr}
	step := 0
	targetLen := len(target)
	for len(queue) > 0 {
		queueSize := len(queue)
		// 开始处理上一轮遗留的元素
		for j := 0; j < queueSize; j++ {
			// 弹出第一个元素（最老的）
			cur := queue[0]
			if len(queue) == 1 {
				queue = []string{}
			} else {
				queue = queue[1:]
			}
			// 必须在这里检查一下，不然如果存在死锁 0000，就会出错，因为后面都是检查过才放入队列的，没办法检查这种情况
			if _, ok := deadMap[cur]; ok {
				continue
			}
			// 分别转动 4 个数字
			for i := 0; i < targetLen; i++ {
				if cur == target {
					return step
				}
				upStr := toUp(cur, i)
				if _, vOk := visited[upStr]; !vOk {
					if _, dOk := deadMap[upStr]; !dOk {
						visited[upStr] = struct{}{}
						queue = append(queue, upStr)
					}
				}
				// 如果不是死亡数字，就放入队列
				downStr := toDown(cur, i)
				if _, vOk := visited[downStr]; !vOk {
					if _, dOk := deadMap[downStr]; !dOk {
						visited[downStr] = struct{}{}
						queue = append(queue, downStr)
					}
				}
			}
		}
		step++
	}

	return -1
}

func getDeadEndsMap(deadends []string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, item := range deadends {
		m[item] = struct{}{}
	}

	return m
}

// 往上
func toUp(lock string, idx int) string {
	status := []byte(lock)
	newNum := status[idx] + 1
	if newNum > '9' {
		status[idx] = '0'
	} else {
		status[idx] = newNum
	}

	return string(status)
}

// 往下
func toDown(lock string, idx int) string {
	status := []byte(lock)
	newNum := status[idx] - 1
	if newNum < '0' {
		status[idx] = '9'
	} else {
		status[idx] = newNum
	}

	return string(status)
}
