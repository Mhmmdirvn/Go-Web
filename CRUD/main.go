package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

type Contact struct {
	Id    int
	Name  string
	Phone string
}

// Func Connect To Database
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Muhammadirvan011206@tcp(127.0.0.1:3306)/db_crud")
	if err != nil {
		return nil, err
	}

	return db, nil
}

var tmpl = template.Must(template.ParseGlob("views/*"))

// Func Read Database
func Index(w http.ResponseWriter, r *http.Request) {
	db, err := Connect()
	rows, err := db.Query("SELECT * FROM tb_crud")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result := []Contact{}

	for rows.Next() {
		each := Contact{}
		err = rows.Scan(&each.Id, &each.Name, &each.Phone)
		if err != nil {
			fmt.Println(err.Error())
		}

		result = append(result, each)
	}

	tmpl.ExecuteTemplate(w, "Index", result)

	defer db.Close()
}

// Func Create
func CreateContact(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "CreateContact", nil)
}

// Func Edit
func EditContact(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "EditContact", nil)
}

// Func Insert Data
func InsertContact(w http.ResponseWriter, r *http.Request) {
	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if r.Method == "POST" {
		name := r.FormValue("name")
		phone := r.FormValue("phone")
		insertForm, err := db.Prepare("INSERT INTO tb_crud (name, phone) VALUES (?, ?)")
		if err != nil {
			fmt.Println(err.Error())
		}

		insertForm.Exec(name, phone)
		log.Println("Insert Data: name " + name + " | phone " + phone)
	}

	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

// Func Edit
func Edit(w http.ResponseWriter, r *http.Request) {
	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	GetId := r.URL.Query().Get("id")
	rows, err := db.Query("SELECT * FROM tb_crud WHERE id=?", GetId)
	if err != nil {
		fmt.Println(err.Error())
	}

	each := Contact{}

	for rows.Next() {
		err = rows.Scan(&each.Id, &each.Name, &each.Phone)
		if err != nil {
			fmt.Println(err.Error())

		}
	}

	tmpl.ExecuteTemplate(w, "EditContact", each)

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
		id := r.FormValue("Getid")
		name := r.FormValue("name")
		phone := r.FormValue("phone")
		insertForm, err := db.Prepare("UPDATE tb_crud SET name=?, phone=? WHERE id=?")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		insertForm.Exec(name, phone, id)
		log.Println("UPDATE: Name: " + name + " | Phone: " + phone)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

// Func Delete
func Delete(w http.ResponseWriter, r *http.Request) {
	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	contactFound := r.URL.Query().Get("id")
	deleteForm, err := db.Prepare("DELETE FROM tb_crud WHERE id=?")
	if err != nil {
		fmt.Println(err.Error())
	}
	deleteForm.Exec(contactFound)
	log.Println("Deleted")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/contact/create", CreateContact)
	http.HandleFunc("/contact/edit", Edit)
	http.HandleFunc("/insert", InsertContact)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))
	
	fmt.Println("serever started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
