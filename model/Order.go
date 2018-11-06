package model

type Order struct {
	Id int64
	User_id int64
	Payment_id int64
	Final_price float64
	Status string
	Deleted_at string
	Created_at string
	Updated_at string
}