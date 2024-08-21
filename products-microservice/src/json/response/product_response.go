package response

type ProductResponse struct {
	ProductID  uint   `json:"product_id"`
	Name       string `json:"name"`
	Category   string `json:"category"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	LastUpdate string `json:"last_update"`
}
