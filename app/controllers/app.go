package controllers

import (
	"github.com/jen6/letzy/app/models"
	"github.com/revel/modules/orm/gorp/app/controllers"
	"github.com/revel/revel"
	//"gopkg.in/Masterminds/squirrel.v1"
)

type App struct {
	gorpController.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) FoodList() revel.Result {
	data := make(map[string]interface{})
	data["error"] = nil

	var foods []models.Food
	_, err := c.Txn.Select(&foods,
		c.Db.SqlStatementBuilder.Select("*").
			From("foods"))

	if err != nil {
		data["error"] = err
	}

	data["foods"] = foods
	return c.RenderJSON(data)
}

func (c App) Rating() revel.Result {
	c.Log.Info("Fetching Info")

	var foods []*models.Food
	_, err := c.Txn.Select(&foods,
		c.Db.SqlStatementBuilder.Select("*").
			From("foods"))

	if err != nil {
		panic(err)
	}
	return c.Render(foods)
}

func (c App) Battle() revel.Result {
	return c.Render()
}

func (c App) MyRating() revel.Result {
	return c.Render()
}
