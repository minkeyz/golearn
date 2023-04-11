package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	x int64
}

type Result struct {
	job    *Job
	result int64
}

var jobChan = make(chan *Job, 100)
var resultChan = make(chan *Result, 100)
var wg sync.WaitGroup

func genInt64(ch chan<- *Job) {
	defer wg.Done()
	for i := 0; i < 60; i++ {
		ix := rand.Int63()
		job := &Job{
			x: ix,
		}
		ch <- job
		time.Sleep(2e4)
	}
	close(ch) // 配合range使用 range会一直等待接受ch 如果不主动关闭 那么就会出现死锁 需要主动关闭
}

func calculator(ch <-chan *Job, res chan<- *Result) {
	defer wg.Done()
	for i := 0; i < 6; i++ {
		ix, ok := <-ch
		if !ok {
			break
		}
		sum := int64(0)
		start := ix.x
		for start > 0 {
			sum += start % 10
			start /= 10
		}
		r := &Result{
			job:    ix,
			result: sum,
		}
		res <- r
		time.Sleep(2e4)
	}
}

func main() {
	wg.Add(1)
	genInt64(jobChan)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go calculator(jobChan, resultChan)
	}
	wg.Wait() // 保证所有 goroutine 执行完毕之后 向下执行

	close(resultChan) // 配合range使用 range会一直等待接受ch 如果不主动关闭 那么就会出现死锁 需要主动关闭
	// fmt.Println(len(resultChan)) // 60
	for v := range resultChan {
		fmt.Printf("%d sum is %d \n", v.job.x, v.result)
	}
}
