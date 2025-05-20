package main

import (
	"sampleProject/bean"
)

func main() {
	employee := new(bean.Employee)
	employee.Id = 1
	employee.Name = "Aashay"
	employee.Salary = 10000
	employee.PrintEmployeeDetail()

	//
	employee1 := bean.Employee{Id: 2, Name: "Mayank", Salary: 10000}

	//

	employee.PrintEmployeeDetail()
	employee1.PrintEmployeeDetail()
}
