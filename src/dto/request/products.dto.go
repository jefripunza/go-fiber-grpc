package request

type CreateProduct struct {
	Code  string `json:"code"`
	Price uint   `json:"price"`
}

type UpdateProduct struct {
	ID    int    `json:"id"`
	Code  string `json:"code"`
	Price uint   `json:"price"`
}
