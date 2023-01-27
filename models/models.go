package models

type Stock struct {
	StockID int64 	`json:"stockid"`
	Name 	string 	`json:"name"`
	Price 	int64 	`json:"price"`
	Company int64 	`json:"company"`
}