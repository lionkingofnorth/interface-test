package main

import "fmt"

type human interface {
	Name() string
	Age() int64
}

type student struct {
	name string
	age  int64
}

func (s student) Name() string {
	return s.name
}

func (s student) Age() int64 {
	return s.age
}

func main() {
	var h human = student{name: "xiaoming", age: 18}
	fmt.Printf("Interface implmented type %T", h)
}
