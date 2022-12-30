package models

import (
	"time"

	"github.com/set2002satoshi/web_contents/api/pkg/module/customs/errors"
	"github.com/set2002satoshi/web_contents/api/pkg/module/types"
)

type ActiveStudentUser struct {
	ActiveStudentUserId types.IDENTIFICATION `gorm:"primaryKey"`
	Class               string               `gorm:"max:4"`
	Name                string               `gorm:"max:16"`
	Wallet              *ActiveWallet        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Login               *ActiveLogin         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Revision            types.REVISION
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func NewActiveStudentUser(
	id int,
	class,
	name string,
	login *ActiveLogin,
	wallet *ActiveWallet,
	revision types.REVISION,
	created_at,
	updated_at time.Time,
) (*ActiveStudentUser, error) {
	au := &ActiveStudentUser{}
	var err error
	err = errors.Combine(err, au.setActiveStudentUserId(id))
	err = errors.Combine(err, au.setClass(class))
	err = errors.Combine(err, au.setName(name))
	err = errors.Combine(err, au.setLogin(login))
	err = errors.Combine(err, au.setWallet(wallet))
	err = errors.Combine(err, au.setRevision(revision))
	err = errors.Combine(err, au.setCreatedAt(created_at))
	err = errors.Combine(err, au.setUpdatedAt(updated_at))
	return au, err
}

func (au *ActiveStudentUser) setActiveStudentUserId(id int) error {
	issuedStudentId, err := types.NewIDENTIFICATION(id)
	if err != nil {
		return err
	}
	au.ActiveStudentUserId = issuedStudentId
	return nil
}

func (au *ActiveStudentUser) setClass(class string) error {
	au.Class = class
	return nil
}

func (au *ActiveStudentUser) setName(name string) error {
	au.Name = name
	return nil
}

func (au *ActiveStudentUser) setLogin(obj *ActiveLogin) error {
	au.Login = obj
	return nil
}

func (au *ActiveStudentUser) setWallet(obj *ActiveWallet) error {
	au.Wallet = obj
	return nil
}

func (au *ActiveStudentUser) setRevision(obj types.REVISION) error {
	au.Revision = obj
	return nil
}

func (au *ActiveStudentUser) setCreatedAt(created_at time.Time) error {
	au.CreatedAt = created_at
	return nil
}

func (au *ActiveStudentUser) setUpdatedAt(updated_at time.Time) error {
	au.UpdatedAt = updated_at
	return nil
}

func (au *ActiveStudentUser) GetActiveStudentUserId() types.IDENTIFICATION {
	return au.ActiveStudentUserId
}

func (au *ActiveStudentUser) GetClass() string {
	return au.Class

}

func (au *ActiveStudentUser) GetName() string {
	return au.Name
}

func (au *ActiveStudentUser) GetLogin() *ActiveLogin {
	return au.Login

}

func (au *ActiveStudentUser) GetWallet() *ActiveWallet {
	return au.Wallet
}

func (au *ActiveStudentUser) GetRevision() types.REVISION {
	return au.Revision
}

func (au *ActiveStudentUser) GetCreatedAt() time.Time {
	return au.CreatedAt
}

func (au *ActiveStudentUser) GetUpdatedAt() time.Time {
	return au.UpdatedAt
}

func (au *ActiveStudentUser) CountUpRevision() error {
	au.Revision += types.REVISION(1)
	return nil
}
