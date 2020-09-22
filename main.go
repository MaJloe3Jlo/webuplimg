package main

import (
	"log"
	"net/http"

	"github.com/MaJloe3Jlo/webuplimg/controllers"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/upload", controllers.UploadFile)
	log.Println("Server started on http://localhost:12356 port")
	http.ListenAndServe(":12356", nil)
}
