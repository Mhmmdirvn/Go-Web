package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/post/create", routeGetindex)
	http.HandleFunc("/post/preview", routeSubmitPost)




	fmt.Println("server started at locahost:9000")
	http.ListenAndServe(":9000", nil)
}

func routeGetindex(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
		var tmpl = template.Must(template.New("form").ParseFiles("inputdata.html"))
		var err = tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func routeSubmitPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tmpl = template.Must(template.New("result").ParseFiles("inputdata.html"))
		
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var title = r.FormValue("title")
		var description = r.Form.Get("description")

		var data = map[string]string{"title": title, "description": description}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}