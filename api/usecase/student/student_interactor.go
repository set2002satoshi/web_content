package student

import (
	"errors"
	"fmt"

	c "github.com/set2002satoshi/web_contents/api/interfaces/controllers"
	"github.com/set2002satoshi/web_contents/api/models"
	customError "github.com/set2002satoshi/web_contents/api/pkg/module/customs/errors"
	"github.com/set2002satoshi/web_contents/api/usecase"
	"github.com/set2002satoshi/web_contents/api/usecase/login"
	"github.com/set2002satoshi/web_contents/api/usecase/wallet"
	"gorm.io/gorm"
)

type StudentInteractor struct {
	DB          usecase.DBRepository
	StudentRepo StudentRepository
	WalletRepo  wallet.WalletRepository
	LoginRepo   login.LoginRepository
}

func (si *StudentInteractor) FindById(ctx c.Context, id int) (*models.ActiveStudentUser, error) {
	db := si.DB.Connect()
	Acquired, err := si.StudentRepo.FindById(db, id)
	if err != nil {
		return nil, err
	}
	return Acquired, nil
}

func (si *StudentInteractor) FindAll(ctx c.Context) ([]*models.ActiveStudentUser, error) {
	db := si.DB.Connect()
	Acquired, err := si.StudentRepo.FindAll(db)
	if err != nil {
		return nil, err
	}
	return Acquired, nil
}

func (si *StudentInteractor) Register(ctx c.Context, obj *models.ActiveStudentUser) (*models.ActiveStudentUser, error) {
	db := si.DB.Connect()
	if !si.isUniqueEmail(db, obj.GetLogin().GetEmail()) {
		return nil, errors.New("このメールアドレスは使われてるます")
	}
	created, err := si.StudentRepo.Create(db, obj)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (si *StudentInteractor) DeleteById(id int) error {
	tx := si.DB.Begin()
	var err error
	fetchedModel, err := si.StudentRepo.FindById(tx, id)
	if err != nil {
		return errors.New("削除対象のstudentが見つかりません")
	}
	fmt.Println(int(fetchedModel.GetLogin().GetActiveLoginId()), int(fetchedModel.GetWallet().GetActiveWalletId()))
	err = customError.Combine(err, si.LoginRepo.DeleteById(tx, fetchedModel))
	err = customError.Combine(err, si.WalletRepo.DeleteById(tx, fetchedModel))
	err = customError.Combine(err, si.StudentRepo.DeleteById(tx, int(fetchedModel.GetActiveStudentUserId())))
	if err != nil {
		tx.Rollback()
		return err
	}
	commitResult := tx.Commit()
	if commitResult.Error != nil {
		return commitResult.Error
	}
	return nil
}

func (si *StudentInteractor) isUniqueEmail(db *gorm.DB, email string) bool {
	count, _ := si.StudentRepo.FetchEmailNumber(db, email)
	return count == 0
}
