package models

type Action struct {
	Idx    int  `db: "idx"`
	FoodA  int  `db: "food_a"`
	FoodB  int  `db: "food_b"`
	IsAWin bool `db: "is_a_win"`
}
