package model

type Books struct {
	Id          string `json:"id" db:"ID"`
	Name        string `json:"name" db:"NAME"`
	Price       string `json:"price" db:"PRICE"`
	Description string `json:"description" db:"DESCRIPTION"`
}

type AllBooks struct {
	Books []Books `json:"books"`
}
