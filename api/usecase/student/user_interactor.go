package student

import (
	"errors"

	c "github.com/set2002satoshi/web_contents/api/interfaces/controllers"
	"github.com/set2002satoshi/web_contents/api/models"
	"github.com/set2002satoshi/web_contents/api/usecase"
	"gorm.io/gorm"
)

type StudentInteractor struct {
	DB      usecase.DBRepository
	Student StudentRepository
}

func (si *StudentInteractor) FindById(ctx c.Context, id int) (*models.ActiveStudentUser, error) {
	db := si.DB.Connect()
	Acquired, err := si.Student.FindById(db, id)
	if err != nil {
		return nil, err
	}
	return Acquired, nil
}

func (si *StudentInteractor) FindAll(ctx c.Context) ([]*models.ActiveStudentUser, error) {
	db := si.DB.Connect()
	Acquired, err := si.Student.FindAll(db)
	if err != nil {
		return nil, err
	}
	return Acquired, nil
}

func (si *StudentInteractor) Register(ctx c.Context, obj *models.ActiveStudentUser) (*models.ActiveStudentUser, error) {
	db := si.DB.Connect()
	if !si.isUniqueEmail(db, obj.GetActiveAuth().GetEmail()) {
		return nil, errors.New("このメールアドレスは使われてるます")
	}
	created, err := si.Student.Create(db, obj)
	if err != nil {
		return nil, err
	}
	return created, nil
}

// シンプルにログは実装しない
func (si *StudentInteractor) Delete(id int) error {
	db := si.DB.Connect()
	err := si.Student.Delete(db, id)
	if err != nil {
		return err
	}
	return nil
}

func (si *StudentInteractor) isUniqueEmail(db *gorm.DB, email string) bool {
	count, _ := si.Student.FetchEmailNumber(db, email)
	return count == 0 
}
