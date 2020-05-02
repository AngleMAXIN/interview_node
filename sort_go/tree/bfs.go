package tree

import "fmt"

var tree map[string][]string

type queue []string

func newQueue(size int) queue {
	return make(queue, size)
}

func (q *queue) pop(idx int) string {
	v := (*q)[idx]
	if idx == 0 {
		*q = (*q)[1:]
	} else if idx > 0 && idx < len(*q)-1 {
		*q = append((*q)[:idx], (*q)[idx+1:]...)
	} else {
		*q = (*q)[:len(*q)-2]
	}
	return v
}

func (q *queue) push(v string) {
	*q = append(*q, v)
}

func (q *queue) len() int {
	return len(*q)
}

// BFS 广度搜索
func BFS(tree map[string][]string, head string) []string {
	queue := newQueue(0)
	mem := make(map[string]struct{})

	queue.push(head)
	mem[head] = struct{}{}
	res := make([]string, 0)
	for queue.len() > 0 {
		node := queue.pop(0)
		nextNodes := tree[node]

		for idx := range nextNodes {
			if _, ok := mem[nextNodes[idx]]; !ok {
				mem[nextNodes[idx]] = struct{}{}
				queue.push(nextNodes[idx])
			}
		}
		res = append(res, node)
	}
	return res
}

