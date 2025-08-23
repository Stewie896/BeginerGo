package main

import (
	"fmt"
	"time"
)

type Person struct {
	name      string
	age       int
	is_adult  bool
	createdAt time.Time
}

type childStruct struct {
	d string
	e string
	f string
}

type parentStruct struct {
	a string
	b string
	c string
	childStruct
}

func CreateUserStruct(name, location, credentials string) any {
	type localscopt struct {
		name     string
		location string
		cred     string
	}

	finalRetruningValue := localscopt{
		name:     name,
		location: location,
		cred:     credentials,
	}
	return &finalRetruningValue
}

func (newId *Person) Mnfnc(newName *string) {
	newId.name = *newName
}

func main() {

	detailes := Person{
		name:      "abhi",
		age:       17,
		is_adult:  false,
		createdAt: time.Now(),
	}

	yalo := Person{
		name:      "yalo",
		age:       19,
		is_adult:  true,
		createdAt: time.Now(),
	}

	changedName := "Npu"

	yalo.Mnfnc(&changedName)

	nm := "abhiral"
	detailes.Mnfnc(&nm)

	fmt.Println(detailes.name)
	fmt.Println(yalo.name)

	child := childStruct{
		d: "Hello iam d form child",
		e: "Hello iam e from child",
		f: "Hello iam f from child",
	}

	fusedDetails := parentStruct{
		a:           "Hello iam a parent",
		b:           "Hello iam b form parent",
		c:           "Hello iam c from parent",
		childStruct: child,
	}

	fmt.Println(fusedDetails, " This is the fused detailes")

	userDetailes := struct {
		name           string
		is_ipv4Allowed bool
	}{
		name:           "User",
		is_ipv4Allowed: true,
	}
	fmt.Println(userDetailes)

	val := CreateUserStruct("abhiral", "llm", "str")
	fmt.Println(val)

}
