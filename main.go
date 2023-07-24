package main

import "fmt"

type person struct {
	firstname   string
	lastname    string
	contactinfo contact
}

type contact struct {
	email string
	zip   int
}

func (p person) printname() {
	fmt.Printf("%+v", p)
}

func main() {
	jim := person{
		firstname:   "jim",
		lastname:    "anderson",
		contactinfo: contact{email: "adanpalma@hotmail.com", zip: 12334},
	}
	jim.printname()

}
