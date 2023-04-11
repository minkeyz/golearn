package main

import "fmt"

func generateNatural() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 2; i < 100000; i++ {
			ch <- i
		}
	}()
	return ch
}

// primeFilter in 单向通道  只能从改通道接受值
func primeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			// 依据上一个 in 和 素数 筛选
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}
func main() {
	ch := generateNatural()
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = primeFilter(ch, prime)
	}
}
