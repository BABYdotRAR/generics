package main

import (
	"fmt"

	"github.com/BABYdotRAR/generics"
	"github.com/BABYdotRAR/generics/data_structures"
)

func main() {
	// Let's create a stack of maps
	stack := data_structures.Stack[map[int]string]{}
	mapA := map[int]string{1: "Rem", 2: "is"}
	mapB := map[int]string{3: "the", 4: "best", 5: "waifu"}
	stack.Push(mapB)
	stack.Push(mapA)
	fmt.Printf("Current stack size: %d\n", stack.Size())
	popA, _ := stack.Pop()
	popB, _ := stack.Pop()
	if stack.IsEmpty() {
		// Now lets use other helpful functions
		_, valuesA := generics.MapKeysAndValues(popA)
		_, valuesB := generics.MapKeysAndValues(popB)
		fmt.Println(valuesA, valuesB)
	}
}
