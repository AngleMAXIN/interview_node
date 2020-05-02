package main

import "fmt"

var mem1 map[int]int

func _numWays(n int) int {
	if n < 2 {
		mem1[n] = 1
		return 1
	}
	v, ok := mem1[n]
	if ok {
		return v
	}
	mem1[n] = _numWays(n-2) + _numWays(n-1)
	return mem1[n]
}

func numWays(n int) int {
	mem1 = make(map[int]int)
	return _numWays(n) % 1000000007
}
func main() {
	n := 36
	res := numWays(n)
	fmt.Println(res)
}
