package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) who() {
	fmt.Println(p.Name)
}

type Writer interface {
	write()
}

type Student struct {
	Person
}

func (s Student) write() {
	fmt.Println(s.Name, "student write")
}

type Teacher struct {
	Person
}

func (t Teacher) write() {
	fmt.Println(t.Name, "teacher write")
}

func main() {
	s := Student{Person{"peter"}}
	s.who()
	s.write()

	t := Teacher{Person{"david"}}
	t.who()
	t.write()
}
