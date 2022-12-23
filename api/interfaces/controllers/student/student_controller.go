package student

import (
	database "github.com/set2002satoshi/web_contents/api/interfaces/database"
	DBstudent "github.com/set2002satoshi/web_contents/api/interfaces/database/student"
	"github.com/set2002satoshi/web_contents/api/models"
	"github.com/set2002satoshi/web_contents/api/pkg/module/dto/response"
	usecase "github.com/set2002satoshi/web_contents/api/usecase/student"
)

type StudentController struct {
	Interactor usecase.StudentInteractor
}

func NewStudentController(db database.DB) *StudentController {
	return &StudentController{
		Interactor: usecase.StudentInteractor{
			DB:      &database.DBRepository{DB: db},
			Student: &DBstudent.StudentRepository{},
		},
	}
}

func (sc *StudentController) convertActiveStudentToDTO(obj *models.ActiveStudentUser) *response.ActiveStudentUserEntity {
	return &response.ActiveStudentUserEntity{
		ActiveStudentUserId: obj.GetActiveStudentUserId(),
		Class:               obj.GetClass(),
		Name:                obj.GetName(),
		Wallet: &models.ActiveWallet{
			ActiveWalletId:      obj.GetWallet().GetActiveWalletId(),
			ActiveStudentUserId: obj.GetWallet().GetActiveStudentUserId(),
			Coin:                obj.GetWallet().GetCoin(),
		},
		ActiveAuth: &models.ActiveStudentAuth{
			Email:               obj.GetActiveAuth().GetEmail(),
			Password:            obj.GetActiveAuth().GetPassword(),
			ActiveStudentUserId: obj.GetActiveAuth().GetStudentId(),
		},
		Option: response.Options{
			Revision:  int(obj.GetRevision()),
			CreatedAt: obj.GetCreatedAt(),
			UpdatedAt: obj.GetUpdatedAt(),
		},
	}
}

func (sc *StudentController) convertActiveStudentToDTOs(obj []*models.ActiveStudentUser) []*response.ActiveStudentUserEntity {
	SEs := make([]*response.ActiveStudentUserEntity, len(obj))
	for i, el := range obj {
		result := &response.ActiveStudentUserEntity{
			ActiveStudentUserId: el.GetActiveStudentUserId(),
			Class:               el.GetClass(),
			Name:                el.GetName(),
			Wallet: &models.ActiveWallet{
				ActiveWalletId:      el.GetWallet().GetActiveWalletId(),
				ActiveStudentUserId: el.GetWallet().GetActiveStudentUserId(),
				Coin:                el.GetWallet().GetCoin(),
			},
			ActiveAuth: &models.ActiveStudentAuth{
				Email:               el.GetActiveAuth().GetEmail(),
				Password:            el.GetActiveAuth().GetPassword(),
				ActiveStudentUserId: el.GetActiveAuth().GetStudentId(),
			},
			Option: response.Options{
				Revision:  int(el.GetRevision()),
				CreatedAt: el.GetCreatedAt(),
				UpdatedAt: el.GetUpdatedAt(),
			},
		}
		SEs[i] = result
	}
	return SEs
}
