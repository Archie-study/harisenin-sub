package main

import (
	"net/http"

	"github.com/Archie-study/harisenin-sub/controllers/bookcontroller"
)

func main() {

	http.HandleFunc("/", bookcontroller.Index)
	http.HandleFunc("/book", bookcontroller.Index)
	http.HandleFunc("/book/index", bookcontroller.Index)
	http.HandleFunc("/book/add", bookcontroller.Add)
	http.HandleFunc("/book/edit", bookcontroller.Edit)
	http.HandleFunc("/book/delete", bookcontroller.Delete)

	http.ListenAndServe(":3001", nil)
}
