package request

type ProductCreateEditModel struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
	ImageUrl string  `json:"imageUrl"`
}
