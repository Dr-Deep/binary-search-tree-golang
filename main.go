package main

import "fmt"

func main() {
	bst := NewBstWrapper([]byte{0})

	fmt.Println(bst)

	bst.Add([]byte{5})
	fmt.Println(bst)

	bst.Add([]byte("hfdhdu"))
	fmt.Println(bst)
}
