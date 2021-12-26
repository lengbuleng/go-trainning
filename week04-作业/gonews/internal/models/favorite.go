package models

type Favorite struct {
	Favoriteid int    `bson:"favoriteid" json:"favoriteid"`
	Newid      string `bson:"newid" json:"newid"`
	Status     bool   `bson:"status"`
}
