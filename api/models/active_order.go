package models

import (
	"github.com/set2002satoshi/web_contents/api/pkg/module/customs/errors"
	"github.com/set2002satoshi/web_contents/api/pkg/module/types"
)

type ActiveOrder struct {
	OrderId          types.IDENTIFICATION `gorm:"primaryKey"`
	ActiveMyTicketId types.IDENTIFICATION
	ActiveFoodId     types.IDENTIFICATION
	NumberOfItem     int
	Price            int
}

func NewActiveOrder(
	orderId int,
	myTicketId int,
	foodId int,
	numberOfItem int,
	price int,
) (*ActiveOrder, error) {
	ao := &ActiveOrder{}
	var err error
	err = errors.Combine(err, ao.setOrderId(orderId))
	err = errors.Combine(err, ao.setMyTicketId(myTicketId))
	err = errors.Combine(err, ao.setFoodId(foodId))
	err = errors.Combine(err, ao.setNumberOfItem(numberOfItem))
	err = errors.Combine(err, ao.setPrice(price))
	return ao, err
}
func (ao *ActiveOrder) GetOrderId() types.IDENTIFICATION {
	return ao.OrderId
}

func (ao *ActiveOrder) GetFoodId() types.IDENTIFICATION {
	return ao.ActiveFoodId
}

func (ao *ActiveOrder) GetNumberOfItem() int {
	return ao.NumberOfItem
}

func (ao *ActiveOrder) GetPrice() int {
	return ao.Price
}

func (ao *ActiveOrder) setOrderId(OrderId int) error {
	issuedOrderId, err := types.NewIDENTIFICATION(OrderId)
	if err != nil {
		return err
	}
	ao.OrderId = issuedOrderId
	return nil
}

func (ao *ActiveOrder) setMyTicketId(OrderId int) error {
	issuedTicketId, err := types.NewIDENTIFICATION(OrderId)
	if err != nil {
		return err
	}
	ao.ActiveMyTicketId = issuedTicketId
	return nil
}

func (ao *ActiveOrder) setFoodId(foodId int) error {
	issuedFoodId, err := types.NewIDENTIFICATION(foodId)
	if err != nil {
		return err
	}
	ao.ActiveFoodId = issuedFoodId
	return nil
}

func (ao *ActiveOrder) setNumberOfItem(NumberOfItem int) error {
	ao.NumberOfItem = NumberOfItem
	return nil
}

func (ao *ActiveOrder) setPrice(Price int) error {
	ao.Price = Price
	return nil
}
