package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var ch = make(chan int)
var wg sync.WaitGroup

func producer() {
	wg.Add(1)
	now := time.Now().UnixNano()
	rand.Seed(now)
	ran := rand.Int()
	ch <- ran
}

func printer() {
	ran := <-ch
	fmt.Println("get it --> " + strconv.Itoa(ran))
	wg.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		time.Sleep(1e6)
		go producer()
	}

	for i := 0; i < 10; i++ {
		go printer()
	}
	wg.Wait()
}
