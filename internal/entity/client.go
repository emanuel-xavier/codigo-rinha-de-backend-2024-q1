package entity

type Client struct {
	Id      int    `json:"id"`
	Name    string `json:"nome"`
	Balance int    `json:"saldo"`
	Limit   int    `json:"limite"`
}
