package main

import "fmt"

func main() {
	r := RandomType{A: "random"}
	q := New()

	q.Insert("test 1")
	fmt.Println(q)
	p1 := q.Remove()
	fmt.Println(q)
	fmt.Println("POPPED: ", p1)
	q.Insert("test 2")
	fmt.Println(q)
	q.Insert(3)
	fmt.Println(q)
	q.Insert(r)
	fmt.Println(q)
	p2 := q.Remove()
	fmt.Println(q)
	fmt.Println("POPPED: ", p2)
	q.Insert("test 5")
	fmt.Println(q)
	q.Insert("test 6")
	fmt.Println(q)
}

type (
	// Queue - An interface for a queue
	Queue interface {
		Insert(interface{})
		Remove()
	}

	// MyQ -
	MyQ []interface{}

	// RandomType -
	RandomType struct {
		A string
	}
)

// New - Create a new MyQ
func New() *MyQ {
	return &MyQ{}
}

// Insert - should insert any type of value into back of MyQ
func (q *MyQ) Insert(we interface{}) {
	*q = append(*q, we)
}

// Remove - should remove value from front of MyQ
func (q *MyQ) Remove() interface{} {

	// You must make a new slice and convert *q values
	s := make([]interface{}, len(*q))
	for i, v := range *q {
		s[i] = v
	}

	popped := s[0]

	*q = append(s[:0], s[1:]...)

	return popped
}
