package student

import (
	c "github.com/set2002satoshi/web_contents/api/interfaces/controllers"
	"github.com/set2002satoshi/web_contents/api/models"
	"github.com/set2002satoshi/web_contents/api/usecase"
)

type StudentInteractor struct {
	DB      usecase.DBRepository
	Student StudentRepository
}

func (si *StudentInteractor) Register(ctx c.Context, obj *models.ActiveStudentUser) (*models.ActiveStudentUser, error) {
	db := si.DB.Connect()
	created, err := si.Student.Create(db, obj)
	if err != nil {
		return nil, err
	}
	return created, nil	
}
