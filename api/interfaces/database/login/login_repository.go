package student

import (
	"github.com/set2002satoshi/web_contents/api/models"
	"gorm.io/gorm"
)

type LoginRepository struct{}

func (repo *LoginRepository) DeleteById(db *gorm.DB, obj *models.ActiveStudentUser) error {
	var login *models.ActiveLogin
	err := db.Unscoped().Delete(&login, obj.GetLogin().GetActiveLoginId())
	if err != nil {
		return err.Error
	}
	return nil
}
