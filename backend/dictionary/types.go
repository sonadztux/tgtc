package dictionary

type Product struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	ShopName     string  `json:"shop_name"`
	ProductPrice float64 `json:"product_price"`
	ImageURL     string  `json:"image_url"`
}

type APIResponseProducts struct {
	Products []Product `json:"products"`
	Error    string    `json:"error_message"`
}

type APIResponseSingleProduct struct {
	Product Product `json:"product"`
	Error   string  `json:"error_message"`
}
