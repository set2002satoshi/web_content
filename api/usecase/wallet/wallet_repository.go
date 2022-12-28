package wallet

import (
	"github.com/set2002satoshi/web_contents/api/models"
	"gorm.io/gorm"
	// "github.com/set2002satoshi/web_contents/api/models"
)

type WalletRepository interface {
	DeleteById(db *gorm.DB, obj *models.ActiveStudentUser) error
}
