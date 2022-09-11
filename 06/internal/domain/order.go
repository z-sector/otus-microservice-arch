package domain

type Order struct {
	ID              int
	ProductID       int
	Quantity        int
	ShippingAddress string
	RequestID       string
}
