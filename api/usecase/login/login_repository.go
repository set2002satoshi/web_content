package login

import (
	"github.com/set2002satoshi/web_contents/api/models"
	"gorm.io/gorm"
)

type LoginRepository interface {
	DeleteById(db *gorm.DB, obj *models.ActiveStudentUser) error
}
