package main

import (
	"fmt"
	"sampleProject/bean"
	"sampleProject/db"
)

func main() {

	// Get Connection from Connection Pool
	db1, err := db.GetConnection()
	if err != nil {
		fmt.Println("Error creating connection pool: ", err.Error())
	}
	defer db1.Close()
	err = db1.Ping()
	if err != nil {
		fmt.Println("Error connecting to the database: ", err.Error())
	} else {
		fmt.Println("Connected to SQL Server successfully!")
	}

	//Performing SQL Operations
	var employeeList []bean.Employee

	employee1 := bean.Employee{1, "Employee1", 1000}
	employee1.PrintEmployeeDetail()
	employee2 := bean.Employee{2, "Employee2", 1000}
	employee2.PrintEmployeeDetail()
	employee3 := bean.Employee{3, "Employee3", 1000}
	employee3.PrintEmployeeDetail()
	employee4 := bean.Employee{4, "Employee4", 1000}
	employee4.PrintEmployeeDetail()

	employeeList = append(employeeList, employee1)
	employeeList = append(employeeList, employee2)
	employeeList = append(employeeList, employee3)
	employeeList = append(employeeList, employee4)

	db.InsertQuery(db1, employeeList)

	employeesList := db.SelectQuery(db1)
	for _, employee := range employeesList {
		fmt.Println(employee)
	}

	db.UpdateQuery(db1, employee4.Id, employee1.Name)

	db.DeleteQuery(db1)
}
