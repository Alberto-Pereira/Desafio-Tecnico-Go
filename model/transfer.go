package model

type Transfer struct {
	ID                     int `json:"id"`
	Account_origin_id      int `json:"account_origin_id"`
	Account_destination_id int `json:"account_destination_id"`
	Amount                 int `json:"amount"`
	Created_at             int `json:"created_at"`
}
