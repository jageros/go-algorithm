package main

import (
	"fmt"
	"sync"
	"time"
)

//func tt(w *sync.WaitGroup, l *sync.Mutex, i int) {
//	defer func() {
//		fmt.Printf("End Goroutine: %d\n", i)
//		l.Unlock()
//		w.Done()
//	}()
//	l.Lock()
//	fmt.Printf("lock: %d\n", i)
//	time.Sleep(time.Second)
//}
//
//func main() {
//	l := &sync.Mutex{}
//	w := &sync.WaitGroup{}
//	for i := 0; i < 10; i++ {
//		w.Add(1)
//		go tt(w, l, i)
//	}
//
//	w.Wait()
//	fmt.Println("Run end!")
//}


func tt(cd *sync.Cond, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	cd.L.Lock()
	fmt.Printf("lock: %d\n", i)
	cd.Wait()
	fmt.Printf("wait end: %d\n", i)
	cd.L.Unlock()
}

func main() {
	l := &sync.Mutex{}
	w := &sync.WaitGroup{}
	cd := sync.NewCond(l)

	for i := 0; i < 10; i++ {
		w.Add(1)
		go tt(cd, w, i)
	}
	time.Sleep(time.Second)
	fmt.Println("Send 5 Signal:")
	for i := 0; i < 5; i++ {
		cd.Signal()   //唤醒一个wait中的goroutine
	}
	time.Sleep(time.Second)
	fmt.Println("Send Broadcast:")
	cd.Broadcast()  //广播唤醒所有wait中的goroutine
	w.Wait()
	fmt.Println("Run end!")
}


//func tt(wg *sync.WaitGroup, i int) {
//	defer wg.Done()
//	fmt.Printf("goroutine: %d\n", i)
//}
//
//func main() {
//	wg := &sync.WaitGroup{}
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go tt(wg, i)
//
//	}
//	wg.Wait()
//	fmt.Println("Run end!")
//}
