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
	StudentId           types.IDENTIFICATION
	Wallet              ActiveWallet
	ActiveAuth          ActiveStudentAuth
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func NewActiveStudentUser(
	id,
	studentId int,
	class,
	name string,
	created_at,
	updated_at time.Time,
) (*ActiveStudentUser, error) {
	au := &ActiveStudentUser{}
	var err error
	err = errors.Combine(err, au.setActiveStudentUserId(id))
	err = errors.Combine(err, au.setActiveStudentId(studentId))
	err = errors.Combine(err, au.setClass(class))
	err = errors.Combine(err, au.setName(name))
	err = errors.Combine(err, au.setCreatedAt(created_at))
	err = errors.Combine(err, au.setUpdatedAt(updated_at))
	return au, err
}

func (au *ActiveStudentUser) setActiveStudentUserId(id int) error {
	issuedStudentId, err := types.NewIDENTIFICATION(id)
	if err != nil {
		return err
	}
	au.StudentId = issuedStudentId
	return nil
}

func (au *ActiveStudentUser) setActiveStudentId(id int) error {
	issuedStudentId, err := types.NewIDENTIFICATION(id)
	if err != nil {
		return err
	}
	au.StudentId = issuedStudentId
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

func (au *ActiveStudentUser) setCreatedAt(created_at time.Time) error {
	au.CreatedAt = created_at
	return nil
}

func (au *ActiveStudentUser) setUpdatedAt(updated_at time.Time) error {
	au.UpdatedAt = updated_at
	return nil
}
