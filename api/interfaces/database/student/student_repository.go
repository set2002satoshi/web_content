package student

import (
	"errors"

	"github.com/set2002satoshi/web_contents/api/models"
	"gorm.io/gorm"
)

type StudentRepository struct{}

func (repo *StudentRepository) Create(db *gorm.DB, obj *models.ActiveStudentUser) (*models.ActiveStudentUser, error) {
	result := db.Create(obj)
	if result.Error != nil {
		return nil, errors.New("ORM層処理が失敗")
	}
	return obj, nil
}
