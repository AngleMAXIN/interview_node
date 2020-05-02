package main

import "fmt"

// var tree map[string][]string

type stack []string

func newStock(size int) stack {
	return make(stack, size)
}

func (q *stack) pop() string {
	if q.len() < 1 {
		return ""
	}
	v := (*q)[q.len()-1]
	*q = (*q)[:q.len()-1]
	return v
}

func (q *stack) push(v string) {
	*q = append(*q, v)
}

func (q *stack) len() int {
	return len(*q)
}

// DFS 深度搜索
func DFS(tree map[string][]string, head string) []string {
	stack := newStock(0)
	mem := make(map[string]struct{})

	stack.push(head)
	mem[head] = struct{}{}
	res := make([]string, 0)
	for stack.len() > 0 {
		node := stack.pop()
		nextNodes := tree[node]

		for i := len(nextNodes) - 1; i >= 0; i-- {
			if _, ok := mem[nextNodes[i]]; !ok {
				fmt.Println(nextNodes[i])
				mem[nextNodes[i]] = struct{}{}
				stack.push(nextNodes[i])
			}
		}
		res = append(res, node)
	}
	return res
}

func main() {
	tree := map[string][]string{
		"2": []string{"3", "4"},
		"3": []string{"5", "6"},
		"4": []string{"7", "8"},
	}
	res := DFS(tree, "2")
	fmt.Println(res)
}
