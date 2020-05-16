package main

import (
	"fmt"
	"github.com/AaronMulgrew/GoFaker"
)

func main() {
	person := faker.GeneratePerson()
	fmt.Printf("%+v\n", person)

	fmt.Println(person.Names.Firstname)
	fmt.Println(person.Names.Surname)
	fmt.Println(person.Bank.IBAN)
	fmt.Println(person.Bank.AccountNumber)
	fmt.Println(person.Bank.BankCode)
	fmt.Println(person.Car.NumberPlate)
	fmt.Println(person.Car.CurrentSpeed)
	
	server := faker.GenerateServer()
	fmt.Printf("%+v\n", server)
	fmt.Println(server.File.Extension)
	fmt.Println(server.File.ContentType)
}
