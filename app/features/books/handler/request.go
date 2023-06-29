package handler 

type RequestCreateBook struct {
	Title         string `json:"title"`
	PublishedYear string `json:"publishedyear"`
	ISBN          string `json:"isbn"` 
	Image         string `json:"image"`
} 

type RequestUpdateBook struct {
	Title         string `json:"title"`
	PublishedYear string `json:"publishedyear"`
	ISBN          string `json:"isbn"`
	Image         string `json:"image"`
} 

