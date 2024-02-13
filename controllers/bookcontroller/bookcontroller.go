package bookcontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/Archie-study/harisenin-sub/entities"
	"github.com/Archie-study/harisenin-sub/libraries"
	"github.com/Archie-study/harisenin-sub/models"
)

var validation = libraries.NewValidation()
var bookModel = models.NewBookModel()

// Handle tampilkan data
func Index(response http.ResponseWriter, request *http.Request) {

	book, _ := bookModel.FindAll()

	data := map[string]interface{}{
		"book": book,
	}

	temp, err := template.ParseFiles("views/book/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

// Handle tambah data
func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {

		temp, err := template.ParseFiles("views/book/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)

	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var book entities.Book
		book.BookTitle = request.Form.Get("book_title")
		book.BookAuthor = request.Form.Get("book_author")
		book.BookYear = request.Form.Get("book_year")
		book.BookDescription = request.Form.Get("book_description")
		book.BookRating = request.Form.Get("book_rating")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(book)

		if vErrors != nil {
			data["book"] = book
			data["validation"] = vErrors
		} else {
			data["message"] = "Book data saved successfully!"
			bookModel.Create(book)
		}

		temp, _ := template.ParseFiles("views/book/add.html")
		temp.Execute(response, data)

	}

}

// Handle edit data
func Edit(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var book entities.Book
		bookModel.Find(id, &book)

		data := map[string]interface{}{
			"book": book,
		}

		temp, err := template.ParseFiles("views/book/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)

	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var book entities.Book
		book.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		book.BookTitle = request.Form.Get("book_title")
		book.BookAuthor = request.Form.Get("book_author")
		book.BookYear = request.Form.Get("book_year")
		book.BookDescription = request.Form.Get("book_description")
		book.BookRating = request.Form.Get("book_rating")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(book)

		if vErrors != nil {
			data["book"] = book
			data["validation"] = vErrors
		} else {
			data["message"] = "Book data has been updated!!!"
			bookModel.Update(book)
		}

		temp, _ := template.ParseFiles("views/book/edit.html")
		temp.Execute(response, data)

	}
}

// Handle delete data
func Delete(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	bookModel.Delete(id)

	http.Redirect(response, request, "/book", http.StatusSeeOther)

}
