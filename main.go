package main

import (
	"log"
	"net/http"
	"study/uplimgweb/controllers"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/upload", controllers.UploadFile)
	log.Println("Server started on :12356 port")
	http.ListenAndServe(":12356", nil)
}
