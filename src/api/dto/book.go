package dto


type CreateBookReq struct {
	Title  string  `json:"title" binding:"required,min=1,max=100"`
	Author string  `json:"author" binding:"required,min=1,max=100"`
	Price  float64 `json:"price" binding:"required,gt=0"`
}

type BookRes struct {
	ID     uint    `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}