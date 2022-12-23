package student

import (
	"gorm.io/gorm"

	"github.com/set2002satoshi/web_contents/api/models"
)

type StudentRepository interface {
	FindById(db *gorm.DB, id int) (*models.ActiveStudentUser, error)
	FindAll(db *gorm.DB) ([]*	models.ActiveStudentUser, error)
	Create(db *gorm.DB, obj *models.ActiveStudentUser) (*models.ActiveStudentUser, error)
}