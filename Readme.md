# Generics! by BABYdotRAR

A collection of some utilities involving generic types.

## Installation

~~~bash
go get github.com/BABYdotRAR/generics
~~~
## Example
~~~go
package main

import (
	"fmt"

	"github.com/BABYdotRAR/generics"
	ds "github.com/BABYdotRAR/generics/data_structures"
)

func main() {
	// Let's create a stack of maps
	stack := ds.Stack[map[int]string]{}
	mapA := map[int]string{1: "Rem", 2: "is"}
	mapB := map[int]string{3: "the", 4: "best", 5: "waifu"}
	stack.Push(mapB)
	stack.Push(mapA)
	fmt.Printf("Current stack size: %d\n", stack.Size())
	popA, _ := stack.Pop()
	popB, _ := stack.Pop()
	if stack.IsEmpty() {
		// Now let's use other helpful functions
		_, valuesA := generics.MapKeysAndValues(popA)
		_, valuesB := generics.MapKeysAndValues(popB)
		fmt.Println(valuesA, valuesB)
	}
}

~~~

### Result
~~~
Current stack size: 2
[Rem is] [the best waifu]
~~~