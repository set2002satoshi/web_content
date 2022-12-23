package student

import (
	"gorm.io/gorm"

	"github.com/set2002satoshi/web_contents/api/models"
)

type StudentRepository interface {
	Create(db *gorm.DB, obj *models.ActiveStudentUser) (*models.ActiveStudentUser, error)
	FindById(db *gorm.DB, id int) (*models.ActiveStudentUser, error)
}