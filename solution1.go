package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int { // HL
	m := map[string]int{}

	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, ".", "", -1)
	sArr := strings.Split(s, " ")
	for _, v := range sArr {
		m[v]++
	}
	return m
} // HL

func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck."
	w := WordCount(s) // HL

	for i, v := range w {
		fmt.Println(i, v)
	}
}
