package main

import (
  "fmt"
  "time"
  "github.com/AaronMulgrew/GoFaker/pkg/data/en_GB/firstnames"
  "github.com/AaronMulgrew/GoFaker/pkg/data/en_GB/automotive"
  "github.com/AaronMulgrew/GoFaker/pkg/data/en_GB/bank"
  "github.com/AaronMulgrew/GoFaker/pkg/data/en_GB/surnames"
)

func main() {
	seed := time.Now().UnixNano()
	firstname := firstnames.GenerateFirstname(seed)
	surname := surnames.GenerateSurname(seed)
	IBAN, AccountNumber, BankCode := bank.GenerateIBAN(seed)
	numberPlate := automotive.GenerateLicensePlate(seed)
	currentSpeed := automotive.GenerateRandomSpeed(seed)
	fmt.Println(firstname)
	fmt.Println(surname)
	fmt.Println(IBAN, AccountNumber, BankCode)
	fmt.Println(numberPlate)
	fmt.Println(currentSpeed)
}