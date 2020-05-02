package main

import (
	"fmt"
)

type testq []int

func (q *testq) append() {

	for i := 0; i < 10; i++ {
		*q = append(*q, i)
	}
	fmt.Println("-", *q)

}

func main() {
	q := make(testq, 0)
	q.append()
	fmt.Println(q)
}
