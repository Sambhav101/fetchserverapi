package receipt

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price string `json:"price"`
}

type Receipt struct {
	Retailer string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total string `json:"total"`
	Items []Item `json:"items"`
}

type ReceiptId struct {
	ID string `json:"id"`
}

type TotalPoints struct {
	TotalPoints int `json:"points"`
}