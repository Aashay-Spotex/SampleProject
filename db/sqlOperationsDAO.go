package db

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"fmt"
	"sampleProject/bean"
)

// returns slice(List) employeeList
func SelectQuery(db1 *sql.DB) []bean.Employee {
	var employeeList []bean.Employee
	query := "SELECT * FROM Employess"
	rows, err := db1.Query(query) // rows is an object similar to ResultSet of java
	if err != nil {
		fmt.Println("Error executing query: ", err.Error())
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows: ", err.Error())
		}
	}(rows)
	for rows.Next() {
		var id, salary int
		var name string
		err = rows.Scan(&id, &name, &salary)
		if err != nil {
			fmt.Println("Error scanning row: ", err.Error())
		} else {
			employeeList = append(employeeList, bean.Employee{
				Name: name, Id: id, Salary: salary,
			})
		}
	}
	return employeeList
}

// takes employeeList as a parameter
func InsertQuery(db1 *sql.DB, employeeList []bean.Employee) {
	query2 := "insert into employess values (@id, @name, @salary)"
	for _, employee := range employeeList {
		_, err := db1.Exec(query2, sql.Named("id", employee.Id), sql.Named("name", employee.Name), sql.Named("salary", employee.Salary))
		if err != nil {
			fmt.Println("Error inserting employess: ", err.Error())
		}
	}
	fmt.Println("Inserted employess successfully!")
}

// string and int parameter
func UpdateQuery(db1 *sql.DB, id int, name string) {
	query1 := "update employess SET name = @name WHERE id = @id"
	_, err := db1.Exec(query1, sql.Named("name", name), sql.Named("id", id))
	if err != nil {
		fmt.Println("Error updating employess: ", err.Error())
	} else {
		fmt.Println("Updated employess successfully!")
	}
}

// no parameter
func DeleteQuery(db1 *sql.DB) {
	query3 := "delete from employess"
	_, err := db1.Exec(query3)
	if err != nil {
		fmt.Println("Error deleting employess: ", err.Error())
	} else {
		fmt.Println("Deleted employess successfully!")
	}
	query3 = "delete from employess_object"
	_, err = db1.Exec(query3)
	if err != nil {
		fmt.Println("Error deleting employess: ", err.Error())
	} else {
		fmt.Println("Deleted employess_object successfully!")
	}
}
func InsertObjectQuery(db1 *sql.DB, employee bean.Employee) {
	gob.Register(bean.Employee{})
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(employee); err != nil {
		fmt.Println("GOB encoding failed: %s", err)
	}
	query4 := "insert into employess_object values (@id, @object)"
	_, err := db1.Exec(query4, sql.Named("id", employee.Id), sql.Named("object", buf.Bytes()))
	if err != nil {
		fmt.Println("Insert failed: %s", err)
	}

	fmt.Println("User stored as BLOB successfully!")
}

func RetrieveObjectQuery(db1 *sql.DB, id int) bean.Employee {
	var blob []byte
	row := db1.QueryRow(`SELECT employee_object FROM employess_object WHERE ID = @p1`, id)
	err := row.Scan(&blob)
	if err != nil {
		fmt.Println("Select failed: %s", err)
	}

	var retrievedEmployee bean.Employee
	decoder := gob.NewDecoder(bytes.NewReader(blob))
	if err := decoder.Decode(&retrievedEmployee); err != nil {
		fmt.Println("GOB decoding failed: %s", err)
	}

	fmt.Printf("Decoded User: %+v\n", retrievedEmployee)
	return retrievedEmployee
}
