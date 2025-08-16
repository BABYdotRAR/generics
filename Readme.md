# Generics! by BABYdotRAR

A collection of some utilities involving generic types.

## Installation

~~~bash
go get github.com/BABYdotRAR/generics@latest
~~~
## Stack Example
~~~go
package main

import (
	"fmt"

	"github.com/BABYdotRAR/generics"
	ds "github.com/BABYdotRAR/generics/data_structures"
)

func stackExample() {
	// Let's create a stack of maps
	stack := ds.NewStack[map[int]string]()
	mapA := map[int]string{1: "Rem", 2: "is"}
	mapB := map[int]string{3: "the", 4: "best", 5: "waifu"}
	stack.Push(mapB)
	stack.Push(mapA)
	fmt.Printf("Current stack size: %d\n", stack.Size())
	popA, _ := stack.Pop()
	popB, _ := stack.Pop()
	if stack.IsEmpty() {
		// Now let's use other helpful functions
		valuesA := generics.MapValues(popA)
		valuesB := generics.MapValues(popB)
		fmt.Println(valuesA, valuesB)
	}
}

~~~

### Result
~~~
Current stack size: 2
[Rem is] [the best waifu]
~~~
## Queue Example
~~~go
package main

import (
	"fmt"

	"github.com/BABYdotRAR/generics"
	ds "github.com/BABYdotRAR/generics/data_structures"
)

func queueExample() {
	// simple queue of numbers
	queue := ds.NewQueue[int]()

	// enqueue some elements
	queue.Enqueue(2)
	queue.Enqueue(4)
	queue.Enqueue(2)
	queue.Enqueue(6)
	queue.Enqueue(5)

	var res []int

	// store the current size of the queue
	res = append(res, queue.Size())
	// store all dequeues
	dequeue, _ := queue.Dequeue()
	res = append(res, dequeue)
	dequeue, _ = queue.Dequeue()
	res = append(res, dequeue)
	dequeue, _ = queue.Dequeue()
	res = append(res, dequeue)
	dequeue, _ = queue.Dequeue()
	res = append(res, dequeue)
	dequeue, _ = queue.Dequeue()
	res = append(res, dequeue)
	// if an empty queue is dequeued, it returns the zero value and false
	zero, ok := queue.Dequeue()
	fmt.Println(zero, ok)
	// let's see what res looks like
	fmt.Println(res)
	// some values are duplicated, let's remove them
	fmt.Println(generics.Unique(res))
}

~~~

### Result
~~~
0 false
[5 2 4 2 6 5]
[5 2 4 6]
~~~