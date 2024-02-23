package entity

import "time"

type Transaction struct {
	Id           int       `json:"id,omitempty"`
	Type         string    `json:"tipo"`
	Descrition   string    `json:"descricao"`
	Value        int       `json:"valor,omitempty"`
	Accomplished time.Time `json:"realizado_em,omitempty"`
	ClientId     int       `json:"id_do_cliente,omitempty"`
}
