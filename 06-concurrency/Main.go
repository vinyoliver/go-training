package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter int = 0
var m = sync.RWMutex{}

func main() {
	fmt.Println("Number of threads ", runtime.GOMAXPROCS(-1)) //number of threads. by default is equal to the number of cores available in the machine
	runtime.GOMAXPROCS(100)                                   //set the number o threads

	//shows where there's some race condition
	//go run -race fileName

	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		sayHello()
		m.Lock()
		increment()
	}
	wg.Wait()

}

func sayHello() {
	fmt.Println("Hello - ", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
