package models

import (
	"github.com/set2002satoshi/web_contents/api/pkg/module/customs/errors"
	"github.com/set2002satoshi/web_contents/api/pkg/module/types"
	"golang.org/x/crypto/bcrypt"
)

type ActiveStudentAuth struct {
	ActiveStudentAuthId   types.IDENTIFICATION `gorm:"primaryKey"`
	Email               string               `gorm:"unique;not null"`
	Password            []byte               `gorm:"not null"`
	ActiveStudentUserId types.IDENTIFICATION
}

func NewActiveStudentAuth(
	id int,
	email string,
	password string,
) (*ActiveStudentAuth, error) {
	as := &ActiveStudentAuth{}
	var err error
	err = errors.Combine(err, as.setStudentId(id))
	err = errors.Combine(err, as.setEmail(email))
	err = errors.Combine(err, as.setPassword(password))
	return as, err
}

func (s *ActiveStudentAuth) setActiveStudentAuthId(id int) error {
	issuedStudentId, err := types.NewIDENTIFICATION(id)
	if err != nil {
		return err
	}
	s.ActiveStudentAuthId = issuedStudentId
	return nil
}

func (s *ActiveStudentAuth) setEmail(email string) error {
	s.Email = email
	return nil
}

func (s *ActiveStudentAuth) setPassword(password string) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	s.Password = []byte(pass)
	return nil
}

func (s *ActiveStudentAuth) setStudentId(id int) error {
	issuedStudentId, err := types.NewIDENTIFICATION(id)
	if err != nil {
		return err
	}
	s.ActiveStudentUserId = issuedStudentId
	return nil
}


func (s *ActiveStudentAuth) GetActiveStudentAuthId() types.IDENTIFICATION {
	return s.ActiveStudentAuthId
}

func (s *ActiveStudentAuth) GetEmail() string {
	return s.Email
}

func (s *ActiveStudentAuth) GetPassword() []byte {
	return s.Password
}

func (s *ActiveStudentAuth) GetStudentId() types.IDENTIFICATION {
	return s.ActiveStudentUserId
}