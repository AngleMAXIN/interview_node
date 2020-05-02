package main

import "fmt"

type node struct {
	val  int
	next *node
}

func deleteNode(head, de *node) {
	p := head
	for p.next != de {
		p = p.next
	}
	p.next = p.next.next
}

func lastRemaining1(n int, m int) int {
	pre := &node{
		val: 0,
	}
	head := pre
	for i := 1; i < n; i++ {
		n := &node{
			val: i,
		}
		pre.next = n
		pre = pre.next
	}
	pre.next = head
	for head.next != head {
		for i := 0; i < m-1; i++ {
			pre = pre.next
			head = head.next
		}
		head = head.next
		pre.next = head
	}
	fmt.Println("--->", head.val)
	return head.val
}

func lastRemaining(n int, m int) int {
	flag := 0
	for i := 2; i <= n; i++ {
		flag = (flag + m) % i
	}
	return flag
}

// public int lastRemaining(int n, int m) {
//     int flag = 0;
//     for (int i = 2; i <= n; i++) {
//         flag = (flag + m) % i;
//         //动态规划的思想，将f(n,m)替换成flag存储
//     }
//     return flag;
// }

func main() {
	lastRemaining(5, 3)
	lastRemaining(5, 1)
	lastRemaining(5, 2)
	lastRemaining(10, 17)
	// fmt.Println("result:", res, res2)
}
