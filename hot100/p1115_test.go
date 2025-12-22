package hot100

import (
	"sync"
	"testing"
)

type FooBar struct {
	n   int
	ch1 chan struct{}
	ch2 chan struct{}
}

func NewFooBar(n int) *FooBar {
	fb := &FooBar{
		n:   n,
		ch1: make(chan struct{}, 1),
		ch2: make(chan struct{}, 1),
	}
	fb.ch1 <- struct{}{}

	return fb
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.ch1
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.ch2 <- struct{}{}
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.ch2
		// printBar() outputs "bar". Do not change or remove this line.
		printBar()

		fb.ch1 <- struct{}{}
	}
}
func TestP1115(t *testing.T) {
	var wg sync.WaitGroup
	fooBar := NewFooBar(100)
	wg.Add(2)
	go func() {
		defer wg.Done()
		fooBar.Foo(func() {
			println("hello")
		})

	}()

	go func() {
		defer wg.Done()
		fooBar.Bar(func() {
			println("bar")
		})
	}()

	wg.Wait()
}
