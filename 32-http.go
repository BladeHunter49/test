package main

import (
	"log"
	"net/http"
	"time"
)

//ver 1
// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("hello,version 1!"))
// 	})
// 	http.HandleFunc("/bye", sayBye)

// 	log.Println("ver1 start")
// 	log.Fatal(http.ListenAndServe(":4000", nil))
// }

// func sayBye(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("bye,version 1!"))
// }

//ver 2
// type myHandle struct{}

// func (_ *myHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("hello,version 2!"))
// }
// func main() {
// 	mux := http.NewServeMux()
// 	mux.Handle("/", &myHandle{})

// 	log.Println("ver2 start")
// 	log.Fatal(http.ListenAndServe(":4000", mux))
// }

//ver 3
type myHandle struct{}

func (_ *myHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello,version 3!"))
}
func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandle{})
	server := &http.Server{
		Addr:         ":4000",
		Handler:      mux,
		WriteTimeout: 2 * time.Second,
	}

	log.Println("ver3 start")
	log.Fatal(server.ListenAndServe())
}
