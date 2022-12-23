package models

import (
	"github.com/set2002satoshi/web_contents/api/pkg/module/customs/errors"
	"github.com/set2002satoshi/web_contents/api/pkg/module/types"
)

type ActiveWallet struct {
	ActiveWalletId      types.IDENTIFICATION `gorm:"primaryKey"`
	ActiveStudentUserId types.IDENTIFICATION
	Coin                uint
}

func NewActiveWallet(
	ActiveWalletId,
	ActiveStudentUserId int,
	coin uint,
) (*ActiveWallet, error) {
	aw := &ActiveWallet{}
	var err error
	err = errors.Combine(err, aw.setActiveWalletId(ActiveWalletId))
	err = errors.Combine(err, aw.setActiveStudentUserId(ActiveStudentUserId))
	err = errors.Combine(err, aw.setCoin(coin))
	return aw, err
}

func (aw *ActiveWallet) setActiveWalletId(id int) error {
	issuedStudentId, err := types.NewIDENTIFICATION(id)
	if err != nil {
		return err
	}
	aw.ActiveWalletId = issuedStudentId
	return nil
}

func (aw *ActiveWallet) setActiveStudentUserId(id int) error {
	issuedStudentId, err := types.NewIDENTIFICATION(id)
	if err != nil {
		return err
	}
	aw.ActiveStudentUserId = issuedStudentId
	return nil
}

func (aw *ActiveWallet) setCoin(coin uint) error {
	aw.Coin += coin
	return nil
}



func (aw *ActiveWallet) GetActiveWalletId() types.IDENTIFICATION {
	return aw.ActiveWalletId
}

func (aw *ActiveWallet) GetActiveStudentUserId() types.IDENTIFICATION {
	return aw.ActiveStudentUserId
}

func (aw *ActiveWallet) GetCoin() uint {
	return aw.Coin
	
}