package bean

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary int
}

// Function of employee Struct
func (employee Employee) PrintEmployeeDetail() {
	fmt.Println("id:", employee.Id, " name:", employee.Name, " salary:", employee.Salary)
}
