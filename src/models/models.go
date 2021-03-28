package models

type Product struct {
	Img         string  `json:"img"`
	Imgalt      string  `json:"imgalt"`
	Price       float64 `json:"price"`
	Promotion   float64 `json:"promotion"`
	ProductName string  `json:"productname"`
	Description string  `json:"desc"`
}

type Customer struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	LoggedIn  bool   `json:"loggedin"`
}

type Order struct {
	Product
	Customer
	CustomerID int `json:"customer_id"`
	ProductID  int `json:"product_id"`
}
