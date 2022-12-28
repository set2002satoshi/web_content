package student

import (
	database "github.com/set2002satoshi/web_contents/api/interfaces/database"
	DBlogin "github.com/set2002satoshi/web_contents/api/interfaces/database/login"
	DBstudent "github.com/set2002satoshi/web_contents/api/interfaces/database/student"
	DBwallet "github.com/set2002satoshi/web_contents/api/interfaces/database/wallet"
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
			DB:          &database.DBRepository{DB: db},
			StudentRepo: &DBstudent.StudentRepository{},
			WalletRepo:  &DBwallet.WalletRepository{},
			LoginRepo:   &DBlogin.LoginRepository{},
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
		Login: &models.ActiveLogin{
			ActiveLoginId:       obj.GetLogin().GetActiveLoginId(),
			ActiveStudentUserId: obj.GetLogin().GetStudentId(),
			Email:               obj.GetLogin().GetEmail(),
			Password:            obj.GetLogin().GetPassword(),
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
			Login: &models.ActiveLogin{
				ActiveLoginId:       el.GetLogin().GetActiveLoginId(),
				ActiveStudentUserId: el.GetLogin().GetStudentId(),
				Email:               el.GetLogin().GetEmail(),
				Password:            el.GetLogin().GetPassword(),
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
