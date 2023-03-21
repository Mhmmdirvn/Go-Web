package main

import(
	"net/http"
	"fmt"
	"html/template"
)

type Info struct {
	Affiliation string
	Address string

}



type Person struct {
	Name string
	Gender string
	Hobbies []string
	Info Info
}
func (t Info) GetAffiliationDetailInfo() string {
	return "Have 31 Divisions"
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Person{
			Name: "Bruce Wayne",
			Gender: "Male",
			Hobbies: []string{"Reading Books", "Traveling", "Buying Things"},
			Info: Info{"Wayne Enterprises", "Ghotam City"},
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}