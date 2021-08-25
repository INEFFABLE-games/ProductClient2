package models

type Product struct {
	Name  string `bson:"name" json:"name"`
	Price uint64 `bson:"price" json:"price"`
}
