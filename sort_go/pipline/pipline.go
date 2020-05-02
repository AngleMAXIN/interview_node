package pipline

import (
	"encoding/binary"
	"io"
	"math/rand"
	"sort"
)

// ArraySource source data
func ArraySource(array ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range array {
			out <- v
		}
		close(out)
	}()
	return out
}

// InMemSort sort nums
func InMemSort(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		a := make([]int, 0)
		for v := range in {
			a = append(a, v)
		}

		sort.Ints(a)

		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

// Merge merge data
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 0)

	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2

		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 < v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
	}()
	return out
}

// ReaderSource source reader
func ReaderSource(reader io.Reader) <-chan int {
	out := make(chan int, 0)
	go func() {
		for {

			buffer := make([]byte, 8)
			n, err := reader.Read(buffer)
			if n > 0 {
				out <- int(binary.BigEndian.Uint64(buffer))
			}
			if err != nil {
				break
			}
		}
		close(out)
	}()

	return out
}

// WriterSink w s
func WriterSink(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

// RandomSource rd
func RandomSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()

	return out
}
