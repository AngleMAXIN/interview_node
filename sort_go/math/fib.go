package main

var mem = make(map[int]int)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	if v, ok := mem[n]; ok {
		return v
	}
	m := fib(n-1) + fib(n-2)
	mem[n] = m % 1000000007
	return mem[n]
}

func fib1(n int) int {
	a := 0
	b := 1
	c := 0
	for i := 0; i < n; i++ {
		c = (a + b) % 1000000007
		a = b
		b = c
	}
	return a % 1000000007
}
