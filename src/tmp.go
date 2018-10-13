package main

import "fmt"

type A struct {
	name string
}

func main (){
	s := 1
	m := A{}
	if s == 1 {
		m = A{"kg"}
	}
	fmt.Println(m)
}