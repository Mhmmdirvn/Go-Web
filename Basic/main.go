// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func handlerIndex(w http.ResponseWriter, r *http.Request) {
// 	var message = "Welcome"
// 	fmt.Fprint(w, message)
// }

// func handlerHello(w http.ResponseWriter, r *http.Request) {
// 	var message = "Hello World!"
// 	fmt.Fprint(w, message)
// }

// // func main() {
// // 	http.HandleFunc("/", handlerIndex)
// // 	http.HandleFunc("/index", handlerIndex)
// // 	http.HandleFunc("/hello", handlerHello)

// // 	var address = ":9000"
// // 	fmt.Printf("server started at %s\n", address)
// // 	err := http.ListenAndServe(address, nil)

// // 	if err != nil {
// // 		fmt.Println(err.Error())
// // 	}
// // }

// func main() {
// 	http.HandleFunc("/", handlerIndex)
// 	http.HandleFunc("/index", handlerIndex)
// 	http.HandleFunc("/hello", handlerHello)

// 	var address = "localhost:9000"
// 	fmt.Printf("server started at %s\n", address)

// 	server := new(http.Server)
// 	server.Addr = address
// 	err := server.ListenAndServe()
// 	if err != nil{
// 		fmt.Println(err.Error())
// 	}
// }

// B.2 Routing http.HandleFunc

package main

import (
	"fmt"
	"net/http"
)	

func main() {
	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	}

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Again")
	})

	fmt.Println("Server Started At localhost:9000")
	http.ListenAndServe("localhost:9000", nil)
}