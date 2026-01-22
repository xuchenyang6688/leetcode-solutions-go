package main

import (
	"fmt"
	"sync"
)

/*
1115. Print FooBar Alternately
*/
type FooBar struct {
	n     int
	fooCh chan struct{}
	barCh chan struct{}
}

func NewFooBar(n int) *FooBar {
	fb := &FooBar{
		n:     n,
		fooCh: make(chan struct{}, 1),
		barCh: make(chan struct{}, 1),
	}
	fb.fooCh <- struct{}{}
	return fb
}

func (fb *FooBar) foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.fooCh
		printFoo()
		fb.barCh <- struct{}{}
	}
}

func (fb *FooBar) bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.barCh
		printBar()
		if i < fb.n-1 {
			fb.fooCh <- struct{}{}
		}
	}
}

func main() {
	foobar := NewFooBar(2)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		foobar.foo(func() { fmt.Print("foo") })
	}()

	go func() {
		defer wg.Done()
		foobar.bar(func() { fmt.Print("bar") })
	}()

	wg.Wait()
	fmt.Println("\nAll completes")

}
