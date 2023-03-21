package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type Student struct {
	Id string
	Name string
	Age int
	Grade int
}

func main() {
	SqlExec()
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Muhammadirvan011206@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}

	return db,nil
}



func SqlQuery() {
	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	var Age = 27
	rows, err := db.Query("select id, name, grade from tb_student where age = ?", Age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Student

	for rows.Next() {
		var each = Student{}
		var err = rows.Scan(&each.Id, &each.Name, &each.Grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)

	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.Name)
	}
}

func SqlQueryRow() {
	var db, err = Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	var result = Student{}
	var id = "E001"
	err = db.
	QueryRow("select name, grade from tb_student where id = ? ", id).
	Scan(&result.Name, &result.Grade)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("name : %s\ngrade: %d\n", result.Name, result.Grade)
}







func SqlPrepare() {
	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	stmt, err := db.Prepare("select name, grade from tb_student where id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result1 = Student{}
	stmt.QueryRow("E001").Scan(&result1.Name, result1.Grade)
	fmt.Printf("name: %s\ngrade: %d\n", result1.Name, result1.Grade)
	
	var result2 = Student{}
	stmt.QueryRow("W001").Scan(&result2.Name, &result2.Grade)
	fmt.Printf("name: %s\ngrade: %d\n", result2.Name, result2.Grade)
	
	var result3 = Student{}
	stmt.QueryRow("B001").Scan(&result3.Name, &result3.Grade)
	fmt.Printf("name: %s\ngrade: %d\n", result3.Name, result3.Grade)
}

func SqlExec() {
	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_,err = db.Exec("insert into tb_student values (?, ?, ?, ?)", "G001", "Galahad", 29, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success!")

	_, err = db.Exec("update tb_student set age = ? where id = ?", 28, "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("update success!")

	_, err = db.Exec("delete from tb_student where id = ?", "G001")
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	fmt.Println("delete success!")

}