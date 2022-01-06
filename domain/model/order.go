package model

type Order struct {
	ID         uint
	OrderType  string
	TargetID   uint
	TargetName string
	Quantity   uint
	User       User
	SubTotal   uint
}
