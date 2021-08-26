package models

type Product struct {
	Name  string `bson:"name" json:"name"`
	Price uint32 `bson:"price" json:"price"`
}
