package model

type Orders struct {
	OrderId     string `json:"orderId" db:"ID"`
	BookDetails Books  `json:"bookDetails"`
	OrderValue  string `json:"orderValue"`
	OrderQty    int    `json:"orderQuantity"`
	Status      string `json:"status" db:"STATUS"`
}
