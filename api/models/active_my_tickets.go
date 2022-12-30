package models

import (
	"time"

	"github.com/set2002satoshi/web_contents/api/pkg/module/customs/errors"
	"github.com/set2002satoshi/web_contents/api/pkg/module/types"
)

type ActiveMyTicket struct {
	TicketId            types.IDENTIFICATION `gorm:"primaryKey"`
	ActiveStudentUserId types.IDENTIFICATION
	Orders              []ActiveOrder
	TotalAmount         int
	DataOfExpiry        time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func NewActiveMyTicket(
	id,
	userId int,
	order []ActiveOrder,
	totalAmount int,
	dataOfExpiry int,
	createdAt,
	updatedAt time.Time,
) (*ActiveMyTicket, error) {
	mt := &ActiveMyTicket{}
	var err error
	err = errors.Combine(err, mt.setTicketId(id))
	err = errors.Combine(err, mt.setUserId(userId))
	err = errors.Combine(err, mt.setOrder(order))
	err = errors.Combine(err, mt.setTotalAmount(totalAmount))
	err = errors.Combine(err, mt.setTotalAmount(dataOfExpiry))
	err = errors.Combine(err, mt.setCreatedAt(createdAt))
	err = errors.Combine(err, mt.setUpdatedAt(updatedAt))

	return mt, err
}

func (amt *ActiveMyTicket) setTicketId(id int) error {
	issuedTicketId, err := types.NewIDENTIFICATION(id)
	if err != nil {
		return err
	}
	amt.TicketId = issuedTicketId
	return nil
}

func (amt *ActiveMyTicket) setUserId(userId int) error {
	issuedUserId, err := types.NewIDENTIFICATION(userId)
	if err != nil {
		return err
	}
	amt.ActiveStudentUserId = issuedUserId
	return nil
}

func (amt *ActiveMyTicket) setOrder(foods []ActiveOrder) error {
	amt.Orders = foods
	return nil
}

func (amt *ActiveMyTicket) setTotalAmount(TotalAmount int) error {
	amt.TotalAmount = TotalAmount
	return nil
}

func (amt *ActiveMyTicket) setCreatedAt(createdAt time.Time) error {
	amt.CreatedAt = createdAt
	return nil
}

func (amt *ActiveMyTicket) setUpdatedAt(updatedAt time.Time) error {
	amt.UpdatedAt = updatedAt
	return nil
}

func (amt *ActiveMyTicket) GetTicketId() types.IDENTIFICATION {
	return amt.TicketId
}

func (amt *ActiveMyTicket) GetUserId() types.IDENTIFICATION {
	return amt.ActiveStudentUserId
}

func (amt *ActiveMyTicket) GetOrders() []ActiveOrder {
	return amt.Orders
}

func (amt *ActiveMyTicket) GetTotalAmount() int {
	return amt.TotalAmount
}

func (amt *ActiveMyTicket) GetDataOfExpiry() time.Time {
	return amt.DataOfExpiry
}

func (amt *ActiveMyTicket) GetCreatedAt() time.Time {
	return amt.CreatedAt
}

func (amt *ActiveMyTicket) GetUpdatedAt() time.Time {
	return amt.UpdatedAt
}
