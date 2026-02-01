package models

type Products struct {
	ID          int32  `json:"id"`
	NAME        string `json:"name"`
	DESCRIPTION string `json:"description"`
	PRICE       int    `json:"price"`
	STOCK       int    `json:"stock"`
}
