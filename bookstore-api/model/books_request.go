package model

type Books struct {
	Id          string `json:"id" db:"ID"`
	Name        string `json:"name" db:"NAME"`
	Price       int    `json:"price" db:"PRICE"`
	Description string `json:"description" db:"DESCRIPTION"`
	Quantity    int    `json:"quantity" db:"QUANTITY"`
}

type AllBooks struct {
	Books []Books `json:"books"`
}
