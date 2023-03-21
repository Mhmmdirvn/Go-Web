package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

type Siswa struct {
	Id      int
	Name    int
	Address int
}

// Mengkoneksikan Ke Database
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Muhammadirvan011206@tcp(127.0.0.1:3306)/db_pendaftaran")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Variable route file
var tmpl = template.Must(template.ParseGlob("views/*"))

// Tampilan Index 
func Index(w http.ResponseWriter, r *http.Request) {
	db, err := Connect()
	rows, err := db.Query("SELECT * FROM tb_siswa")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result := []Siswa{}

	for rows.Next() {
		each := Siswa{}
		err = rows.Scan(&each.Id, &each.Name, &each.Address)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, each)
	}

	tmpl.ExecuteTemplate(w, "Index", result)

	defer db.Close()
}

// Add
func Add(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Add", nil)
}

// Edit Form
func EditForm(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "EditForm", nil)
}

// Add New Data
func AddNewdata(w http.ResponseWriter, r *http.Request) {
	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if r.Method == "POST" {
		name := r.FormValue("name")
		address := r.FormValue("address")
		ins, err := db.Prepare("INSERT INTO tb_siswa (name, address) VALUES (?, ?)")
		if err != nil {
			panic(err.Error())
		}
		ins.Exec(name, address)
		log.Println("INSERT DATA: name" + name + " | address " + address)
	}

	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

// Edit 
func Edit(w http.ResponseWriter, r *http.Request) {
	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	getId := r.URL.Query().Get("id")
	rows, err := db.Query("SELECT * FROM tb_siswa WHERE id=?", getId)
	if err != nil {
		fmt.Println(err.Error())
	}

	each := Siswa{}

	for rows.Next() {
		err = rows.Scan(&each.Id, &each.Name, &each.Address)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	tmpl.ExecuteTemplate(w, "EditForm", each)

	defer db.Close()
}

// Func Update 
func Update(w http.ResponseWriter, r *http.Request) {
	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if r.Method == "POST" {
		id := r.FormValue("getId")
		name := r.FormValue("name")
		address := r.FormValue("address")
		insertform, err := db.Prepare("UPDATE tb_siswa SET name=?, address=? WHERE id=?")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		insertform.Exec(name, address, id)
		log.Println("UPDATE: Name: " + name + " | Address :" + address)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

// Delete
func Delete(w http.ResponseWriter, r *http.Request) {
	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	getData := r.URL.Query().Get("id")
	delete, err := db.Prepare("DELETE FROM tb_siswa WHERE id=?")
	if err != nil {
		fmt.Println(err.Error())
	}

	delete.Exec(getData)
	log.Println("Deleted")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/create", Add)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", AddNewdata)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
