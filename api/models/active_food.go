package models

import (
	"github.com/set2002satoshi/web_contents/api/pkg/module/customs/errors"
	"github.com/set2002satoshi/web_contents/api/pkg/module/types"
)

type ActiveFood struct {
	FoodId       types.IDENTIFICATION `gorm:"primaryKey"`
	Order        []ActiveOrder
	Name         string
	FoodImageURL string
	Price        int
}

func NewActiveFood(
	FoodId int,
	order []ActiveOrder,
	name string,
	foodImageURL string,
	price int,
) (*ActiveFood, error) {
	af := &ActiveFood{}
	var err error
	err = errors.Combine(err, af.setActiveFoodId(FoodId))
	err = errors.Combine(err, af.setActiveOrder(order))
	err = errors.Combine(err, af.setName(name))
	err = errors.Combine(err, af.setFoodImageURL(foodImageURL))
	err = errors.Combine(err, af.setPrice(price))
	return af, err
}

func (af *ActiveFood) setActiveFoodId(id int) error {
	issuedFoodId, err := types.NewIDENTIFICATION(id)
	if err != nil {
		return err
	}
	af.FoodId = issuedFoodId
	return nil
}

func (af *ActiveFood) setActiveOrder(order []ActiveOrder) error {
	af.Order = order
	return nil
}

func (af *ActiveFood) setName(name string) error {
	af.Name = name
	return nil
}

func (af *ActiveFood) setFoodImageURL(image string) error {
	af.FoodImageURL = image
	return nil
}
func (af *ActiveFood) setPrice(price int) error {
	af.Price = price
	return nil
}
