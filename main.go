package main

import (
	"fmt"
	"github.com/Abishnoi69/WebRouter/modules"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/hello", modules.HelloHandler)
	http.HandleFunc("/form", modules.FormHandler)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
