package handler

type ResponseGetBooks struct {
	Title         string `json:"title"`
	PublishedYear string `json:"publishedyear"`
	ISBN          string `json:"isbn"` 
	Image         string `json:"image"`
} 

type ResponseGetBook struct {
	Title         string `json:"title"`
	PublishedYear string `json:"publishedyear"`
	ISBN          string `json:"isbn"` 
	Image         string `json:"image"`
} 


type BookResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    BookData  `json:"data"`
}

type BookData struct {
	Title         string `json:"title"`
	PublishedYear string `json:"publishedyear"`
	ISBN          string `json:"isbn"` 
	Image         string `json:"image"`
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
	Title         string `json:"title"`
	PublishedYear string `json:"publishedyear"`
	ISBN          string `json:"isbn"` 
	Image         string `json:"image"`
}
