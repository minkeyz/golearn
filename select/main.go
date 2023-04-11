package main

import (
	"fmt"
	"strconv"
	"sync"
)

//func main() {
//	ch := make(chan int, 1)
//	for i := 0; i < 20; i++ {
//		select {
//		case x := <-ch:
//			fmt.Println("x -> ", x)
//		case y := <-ch:
//			fmt.Println("y -> ", y)
//		default:
//			fmt.Println(i)
//			ch <- i
//		}
//	}
//}

//var wg sync.WaitGroup
//
//func printer(s string) {
//	fmt.Printf("a log test -> %s \n", s)
//}
//
//func worker(ch chan<- string, start int) {
//	defer wg.Done()
//	for i := 100 * start; i < 100*(start+1); i++ {
//		ix := i%100 == 0
//		if ix {
//			// 报错了
//			ch <- strconv.Itoa(i)
//		}
//	}
//}
//
//// 模仿 日志打印耗时操作
//func watcher(ch <-chan string) {
//	for {
//		time.Sleep(2e9)
//		v, ok := <-ch
//		if !ok {
//			fmt.Println("done")
//			wg.Done()
//		}
//		printer(v)
//	}
//}
//
//func main() {
//	ch := make(chan string, 1000)
//	wg.Add(10)
//	for i := 0; i < 10; i++ {
//		go worker(ch, i)
//	}
//	wg.Wait()
//	close(ch)
//	wg.Add(1)
//	go watcher(ch)
//	wg.Wait()
//}

var wg sync.WaitGroup

func printer(s string) {
	fmt.Printf("a log test -> %s \n", s)
}

// 生产者
func worker(ch chan<- string) {
	for i := 1; i < 1000; i++ {
		ix := i%100 == 0
		if ix {
			// 报错了
			ch <- strconv.Itoa(i)
			fmt.Printf("wrong  %d \n", i)
		}
	}
	// 主程序结束时候 当然可以不结束 --> for
	close(ch)
}

// 模仿 日志打印耗时操作 消费者
func watcher(ch <-chan string) {
	for {
		v, ok := <-ch
		if !ok {
			// 通道关闭时候 触发
			fmt.Println("done")
			wg.Done()
			break
		}
		printer(v)
	}
}

func needLog(f func(chan<- string), ch chan string) {
	wg.Add(1)
	go watcher(ch)
	f(ch)
	wg.Wait()
}

func main() {
	ch := make(chan string, 1000)
	needLog(worker, ch)
}
