package handler 

type RequestCreateBook struct {
	Title         string `json:"title"`
	PublishedYear string `from:"publishedyear"`
	ISBN          string `from:"isbn"`
} 

type RequestUpdateBook struct {
	Title         string `from:"title"`
	PublishedYear string `from:"publishedyear"`
	ISBN          string `from:"isbn"`
} 

