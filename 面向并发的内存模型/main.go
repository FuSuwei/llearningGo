package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total struct {
	sync.Mutex
	value int
}
var total2 uint64


func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 100; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

// "sync/atomic" 内置方法
func worker2(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for i = 0; i <= 100; i++ {
		atomic.AddUint64(&total2, i)
	}
}



type singleton struct {}

var (
	instance *singleton
	once     sync.Once
)

func Instance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}



func main() {
	//var wg sync.WaitGroup
	//wg.Add(2)
	////go worker(&wg)
	////go worker(&wg)
	//go worker2(&wg)
	//go worker2(&wg)
	//wg.Wait()
	//fmt.Println(total2)

	//var a int64 = 1
	//atomic.SwapInt64(&a, 23)
	//fmt.Println(a)

	a := Instance()
	fmt.Println(a)
}

func sum(){
	s := 0
	for i := 0; i <= 100; i++ {
		s += i
	}
	fmt.Println(s)
}