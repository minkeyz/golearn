package main

import (
	"fmt"
	"time"
)

var myMap = make(map[int]uint64, 0)

func Dg(n int) uint64 {
	if n == 1 || n == 2 {
		return 1
	} else {
		if myMap[n] > 0 {
			return myMap[n]
		} else {
			myMap[n-1] = Dg(n - 1)
			myMap[n-2] = Dg(n - 2)
			return myMap[n-1] + myMap[n-2]
		}
	}
}

func Fb(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return Fb(n-1) + Fb(n-2)
}

func main() {
	start := time.Now().UnixNano()
	dg := Dg(100)
	fmt.Println(dg)
	end := time.Now().UnixNano()
	fmt.Println("time ->", end-start)
}
