package p677

type MapSum struct {
	stage map[string]int
	pre   map[string]int
}

func Constructor() MapSum {
	return MapSum{stage: make(map[string]int), pre: make(map[string]int)}
}

func (m *MapSum) Insert(key string, val int) {
	if oldVal, ok := m.stage[key]; ok {
		if oldVal == val {
			return
		}
		for i := 1; i <= len(key); i++ {
			m.pre[key[0:i]] += (val - oldVal)
		}
		m.stage[key] = val
		return
	}
	for i := 1; i <= len(key); i++ {
		m.pre[key[0:i]] += val
	}
	m.stage[key] = val
}

func (m *MapSum) Sum(prefix string) int {
	return m.pre[prefix]
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
