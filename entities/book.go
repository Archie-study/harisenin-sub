package entities

type Book struct {
	Id              int64
	BookTitle       string `validate:"required" label:"Book Title"`
	BookAuthor      string
	BookYear        string
	BookDescription string
	BookRating      string `validate:"required" label:"Book Rating"`
}
