package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	pbuilder := &PersonBuilder{}
	pp, err := pbuilder.SetName("Paul", "").SetAge(165).Build()
	if err != nil {
		return err
	}
	fmt.Println(pp)
	return nil
}

//Person rappersent a human
type Person struct {
	//Name is how person is called
	Name string
	//Surname is how person is identify
	Surname string
	//time pass by
	Age int
}

//PersonBuilder for buildint a person object safely
type PersonBuilder struct {
	p Person
	e error
}

//Build is final act of PersonBuilder
func (pb *PersonBuilder) Build() (Person, error) {
	return pb.p, pb.e
}

//SetName check and set person name and surname
func (pb *PersonBuilder) SetName(name, surname string) *PersonBuilder {
	if name == "" || surname == "" {
		pb.e = manageError(pb.e, "no valid name")
	}

	pb.p.Name = name
	pb.p.Surname = surname
	return pb
}

//SetAge check and set person age
func (pb *PersonBuilder) SetAge(age int) *PersonBuilder {
	if age <= 0 || age > 101 {
		pb.e = manageError(pb.e, "no valid age")
		return pb
	}

	pb.p.Age = age
	return pb
}

func manageError(prevError error, message string) error {
	var e error
	if prevError != nil {
		e = errors.New(prevError.Error() + " and " + message)
	} else {
		e = errors.New(message)
	}
	return e
}
