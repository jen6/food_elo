package models

type Food struct {
	Idx    int    `db: "idx"`
	Name   string `db: "name"`
	Rating int    `db: "rating"`
}
