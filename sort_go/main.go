package main

import (
	"bufio"
	"fmt"
	"os"
	"srot_go/pipline"
	"srot_go/sort"
)

const (
	// N slice lenght
	N = 40
	// Limit slice values range
	Limit = 120
)

func main1() {
	s := []int{4, 5, 6}
	n := copy(s, s[1:])
	fmt.Println(len(s), s, n)
}

func main2() {
	// rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, N)
	// for i := 0; i < N; i++ {
	// 	nums = append(nums, rand.Intn(Limit))
	// }
	nums = []int{1, 2, 3, 4, 4, 5, 6, 8}
	fmt.Println("before:", nums)
	sort.QuickSort(nums)
	// sort.FindSort(nums)
	// sort.MergeSort(nums)
	fmt.Println("after:", nums)

	// for i := 1; i < N; i++ {
	// 	if nums[i] < nums[i-1] {
	// 		fmt.Println("error", nums[i], nums[i-1])
	// 	}
	// }
}

func mergeDemo() {
	p := pipline.Merge(
		pipline.InMemSort(pipline.ArraySource(3, 5, 8, 1, 34, 7, 0, 56)),
		pipline.InMemSort(pipline.ArraySource(6, 53, 8, 134, 3444, 73, 10, 256)),
	)

	for v := range p {
		fmt.Println(v)
	}
}

func main3() {
	const filename = "small.in"
	const count = 10000000
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	p := pipline.RandomSource(count)
	writer := bufio.NewWriter(f)

	pipline.WriterSink(writer, p)
	writer.Flush()

	f, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	out := pipline.ReaderSource(bufio.NewReader(f))
	for v := range out {
		fmt.Println(v)
	}

}

func main() {
	res := de()
	fmt.Println(res)
}

func de() int {
	a := 54
	defer func() {
		fmt.Printf("%p", &a)
		a = 1
	}()
	a = 34
	fmt.Printf("1---%p\n", &a)

	return a
}
