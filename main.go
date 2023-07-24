package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p person) updateName(newName string) {
	p.firstname = newName
}

func main() {
	jim := person{
		firstname: "jim",
		lastname:  "anderson",
		contactInfo: contactInfo{email: "adanpalma@hotmail.com",
			zip: 12334},
	}
	jim.print()
	jim.updateName("Jimmy")
	jim.print()

	a := &jim
	fmt.Printf("%v su valor es %v", &a, *(&jim))

}
