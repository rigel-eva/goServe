package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

const port = ":8080"

func serveHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	fmt.Printf("%v : Page Served! %v\n", time.Now(), r.URL.Path)
	w.Write([]byte(message))
}
func main() {
	http.HandleFunc("/", serveHello)
	fmt.Printf("%v : Server started on %v", time.Now(), port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
	fmt.Println("Exited, see you soon!")
}
