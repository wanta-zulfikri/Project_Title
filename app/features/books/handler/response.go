package handler

type ResponseGetBooks struct {
	Title         string `from:"title"`
	PublishedYear string `from:"publishedyear"`
	ISBN          string `from:"isbn"`
} 

type ResponseGetBook struct {
	Title         string `from:"title"`
	PublishedYear string `from:"publishedyear"`
	ISBN          string `from:"isbn"`
} 


type BookResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    BookData  `json:"data"`
}

type BookData struct {
	Title         string `from:"title"`
	PublishedYear string `from:"publishedyear"`
	ISBN          string `from:"isbn"`
}



type BooksResponse struct {
	Code       int                 `json:"code"`
	Message    string              `json:"message"`
	Data       []ResponseGetBooks  `json:"data"`
	Pagination Pagination          `json:"pagination"`
}

type Pagination struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
}

type ResponseUpdateBooks struct {
	Title         string `from:"title"`
	PublishedYear string `from:"publishedyear"`
	ISBN          string `from:"isbn"`
}
