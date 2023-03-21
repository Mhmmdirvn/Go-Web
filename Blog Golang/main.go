package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type LinkPost struct {
	Post string
	Link string
}

type Index struct {
	Text string
	Link []LinkPost
}

type Post struct {
	Title string
	Description string
	ImgSrc      string
	Comments []ListComment
}

type ListComment struct {
	Name string
	Comment string
}

func main() {
	// Index
	var tmpl, err = template.ParseGlob("views/*")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var index = Index{
			Link: []LinkPost{
				{Post: "Judul Posting 1", Link: "/post/1"},
				{Post: "Judul Posting 2", Link: "/post/2"},
				{Post: "Judul Posting 3", Link: "/post/3"},
			},
			Text: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Alias rem dolore minima, ipsum voluptates, dolorem, deserunt doloremque mollitia harum commodi ut assumenda reprehenderit! Accusamus dignissimos officia unde harum animi veritatis.",
		}

		err = tmpl.ExecuteTemplate(w, "index", index)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Post 1
	http.HandleFunc("/post/1", func(w http.ResponseWriter, r *http.Request) {
		var post1 = Post{
			ImgSrc:      "img/img.png",
			Description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Voluptatem deserunt nobis tenetur nulla ipsum beatae, vitae quibusdam iste rerum cupiditate fuga amet, quod accusamus magni laborum pariatur vero, dolores voluptate!",
			Comments: []ListComment{
				{Name: "M.Nindra Zaka", Comment: "Artikelnya Bagus, Mantap"},
				{Name: "Ilham Surya", Comment: "Keren Sekali"},
				{Name: "Yaza", Comment: "Sangat Membantu"},
			},
			Title: "Judul Posting 1",
		}

		err = tmpl.ExecuteTemplate(w, "post1", post1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Post 2
	http.HandleFunc("/post/2", func(w http.ResponseWriter, r *http.Request) {
		var post2 = Post{
			ImgSrc:      "img/img.png",
			Description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Voluptatem deserunt nobis tenetur nulla ipsum beatae, vitae quibusdam iste rerum cupiditate fuga amet, quod accusamus magni laborum pariatur vero, dolores voluptate!",
			Comments: []ListComment{
				{Name: "Irvan", Comment: "Gokil"},
				{Name: "Alif", Comment: "Mantap Sekali Gan"},
				{Name: "Yaza", Comment: "Sangat Membantu"},
			},
			Title: "Judul Posting 2",
		}

		err = tmpl.ExecuteTemplate(w, "post1", post2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Post 3
	http.HandleFunc("/post/3", func(w http.ResponseWriter, r *http.Request) {
		var post3 = Post{
			ImgSrc:      "img/img.png",
			Description: "Lorem ipsum, dolor sit amet consectetur adipisicing elit. Voluptatem deserunt nobis tenetur nulla ipsum beatae, vitae quibusdam iste rerum cupiditate fuga amet, quod accusamus magni laborum pariatur vero, dolores voluptate!",
			Comments: []ListComment{
				{Name: "Ilul", Comment: "Aku Suka"},
				{Name: "Aka", Comment: "Keren Sekali"},
				{Name: "Doni", Comment: "Keren Sih"},
			},
			Title: "Judul Posting 3",
		}

		err = tmpl.ExecuteTemplate(w, "post1", post3)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Input Data
	http.HandleFunc("/post/create", routeGetindex)
	http.HandleFunc("/post/preview", routeSubmitPost)


	// Routing
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("img"))))

	// Server
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func routeGetindex(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
		var tmpl = template.Must(template.New("form").ParseFiles("views/inputdata.html"))
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
		var tmpl = template.Must(template.New("result").ParseFiles("views/inputdata.html"))
		
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