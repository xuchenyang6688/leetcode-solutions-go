package main // Changed from print_in_order

import (
	"fmt"
	"sync"
)

type Foo struct {
	firstDone  chan struct{}
	secondDone chan struct{}
}

func NewFoo() *Foo {
	return &Foo{
		firstDone:  make(chan struct{}),
		secondDone: make(chan struct{}),
	}
}

func (f *Foo) First(printFirst func()) {
	printFirst()
	close(f.firstDone)
}

func (f *Foo) Second(printSecond func()) {
	<-f.firstDone
	printSecond()
	close(f.secondDone)
}

func (f *Foo) Third(printThird func()) {
	<-f.secondDone
	printThird()
}

func main() { // Added main function
	foo := NewFoo()
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		foo.Third(func() { fmt.Print("third") })
	}()

	go func() {
		defer wg.Done()
		foo.First(func() { fmt.Print("first") })
	}()

	go func() {
		defer wg.Done()
		foo.Second(func() { fmt.Print("second") })
	}()

	wg.Wait()
	fmt.Println("\nAll completes")
}
