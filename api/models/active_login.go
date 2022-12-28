package models

import (
	"github.com/set2002satoshi/web_contents/api/pkg/module/customs/errors"
	"github.com/set2002satoshi/web_contents/api/pkg/module/types"
	"golang.org/x/crypto/bcrypt"
)

type ActiveLogin struct {
	ActiveLoginId types.IDENTIFICATION `gorm:"primaryKey"`
	ActiveStudentUserId types.IDENTIFICATION
	Email               string `gorm:"unique;not null"`
	Password            []byte `gorm:"not null"`
}

func NewActiveLogin(
	id int,
	studentId int,
	email string,
	password string,
) (*ActiveLogin, error) {
	as := &ActiveLogin{}
	var err error
	err = errors.Combine(err, as.setActiveLoginId(id))
	err = errors.Combine(err, as.setStudentId(studentId))
	err = errors.Combine(err, as.setEmail(email))
	err = errors.Combine(err, as.setPassword(password))
	return as, err
}

func (s *ActiveLogin) setActiveLoginId(id int) error {
	issuedStudentId, err := types.NewIDENTIFICATION(id)
	if err != nil {
		return err
	}
	s.ActiveLoginId = issuedStudentId
	return nil
}

func (s *ActiveLogin) setEmail(email string) error {
	s.Email = email
	return nil
}

func (s *ActiveLogin) setPassword(password string) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	s.Password = []byte(pass)
	return nil
}

func (s *ActiveLogin) setStudentId(id int) error {
	issuedStudentId, err := types.NewIDENTIFICATION(id)
	if err != nil {
		return err
	}
	s.ActiveStudentUserId = issuedStudentId
	return nil
}

func (s *ActiveLogin) GetActiveLoginId() types.IDENTIFICATION {
	return s.ActiveLoginId
}

func (s *ActiveLogin) GetEmail() string {
	return s.Email
}

func (s *ActiveLogin) GetPassword() []byte {
	return s.Password
}

func (s *ActiveLogin) GetStudentId() types.IDENTIFICATION {
	return s.ActiveStudentUserId
}
