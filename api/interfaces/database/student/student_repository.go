package student

import (
	"errors"

	"github.com/set2002satoshi/web_contents/api/models"
	"gorm.io/gorm"
)

type StudentRepository struct{}

func (repo *StudentRepository) FindById(db *gorm.DB, id int) (*models.ActiveStudentUser, error) {
	var student *models.ActiveStudentUser
	result := db.Where("active_student_user_id = ?", id).Preload("Wallet").Preload("Login").Find(&student)
	if result.Error != nil {
		return nil, result.Error
	}
	return student, nil
}

func (repo *StudentRepository) FetchEmailNumber(db *gorm.DB, email string) (int64, error) {
	var asa *models.ActiveLogin
	var count int64
	db.Model(&asa).Where("email = ?", email).Count(&count)
	return count, nil
}

func (repo *StudentRepository) FindAll(db *gorm.DB) ([]*models.ActiveStudentUser, error) {
	var student []*models.ActiveStudentUser
	result := db.Preload("Wallet").Preload("Login").Find(&student)
	if result.Error != nil {
		return nil, result.Error
	}
	return student, nil
}

func (repo *StudentRepository) Create(db *gorm.DB, obj *models.ActiveStudentUser) (*models.ActiveStudentUser, error) {
	result := db.Create(&obj)
	if result.Error != nil {
		return nil, errors.New("ORM層処理が失敗")
	}
	return obj, nil
}

func (repo *StudentRepository) DeleteById(db *gorm.DB, id int) error {
	var student *models.ActiveStudentUser
	result := db.Unscoped().Delete(&student, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
