package student

import (
	"github.com/set2002satoshi/web_contents/api/models"
	"gorm.io/gorm"
)

type WalletRepository struct{}

func (repo *WalletRepository) DeleteById(db *gorm.DB, obj *models.ActiveStudentUser) error {
	var wallet *models.ActiveWallet
	err := db.Unscoped().Delete(&wallet, obj.GetWallet().GetActiveWalletId())
	if err != nil {
		return err.Error
	}
	return nil
}
