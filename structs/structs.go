package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Lee"
	// fmt.Println(alex)
	// fmt.Printf("%+v\n", alex)
	jim := person{
		firstName: "Jim",
		lastName:  "Kim",
		contactInfo: contactInfo{
			email:   "Jim@gmail.com",
			zipCode: 94000,
		},
	}
	// jimPointer := &jim // give me the memory address of jim
	jim.updateName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName // (*p) = give me the value of p , p is a pointer (memory address) to jim, so get me the value of p by
	// dereferencing it by *p
}
