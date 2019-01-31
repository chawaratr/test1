package main

import "fmt"
import (
	"kbtg.training/class2/hw/pkg"
)

func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck."
	fmt.Println(s)
	w := words.WordCount(s) // HL

	for i, v := range w {
		fmt.Println(i, v)
	}
}
