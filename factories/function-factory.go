package main

import (
	"fmt"
)

type Employee struct {
	Name string
	Position string
	AnnualIncome float64
}

func NewEmployeeFactory(position string,annualIncome float64) func(name string ) *Employee{
	return func(name string) *Employee{
      return &Employee{Name: name,Position: position,AnnualIncome: annualIncome}
	}
}

func main() {
   developerFactory:=NewEmployeeFactory("develop",6500)
   managerFactory:=NewEmployeeFactory("manager",16000)

   developer:=developerFactory("Adam")
   manager:=managerFactory("jane")

   fmt.Println(developer)
   fmt.Println(manager)

}
