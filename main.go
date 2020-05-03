package main

import (
	"fmt"
	"goFaker"
)

func main() {
	person := faker.GeneratePerson()

	fmt.Println(person.Names.Firstname)
	fmt.Println(person.Names.Surname)
	fmt.Println(person.Bank.IBAN)
	fmt.Println(person.Bank.AccountNumber)
	fmt.Println(person.Bank.BankCode)
	fmt.Println(person.Car.NumberPlate)
	fmt.Println(person.Car.CurrentSpeed)
}
