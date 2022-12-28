package response

import (
	"github.com/set2002satoshi/web_contents/api/models"
	"github.com/set2002satoshi/web_contents/api/pkg/module/types"
)

type (
	FindAllActiveStudentUserResponse struct {
		Result *ActiveStudentUserResults `json:"result"`

		CodeErr error  `json:"error"`
		MsgErr  string `json:"msg"`
	}

	FindByIDActiveStudentUserResponse struct {
		Result *ActiveStudentUserResult `json:"result"`

		CodeErr error  `json:"error"`
		MsgErr  string `json:"msg"`
	}

	CreateActiveStudentUserResponse struct {
		Result *ActiveStudentUserResult `json:"results"`

		CodeErr error  `json:"code"`
		MsgErr  string `json:"msg"`
	}
)

type (
	ActiveStudentUserResult struct {
		Student *ActiveStudentUserEntity `json:"student"`
	}
	ActiveStudentUserResults struct {
		Students []*ActiveStudentUserEntity `json:"students"`
	}

	HistoryStudentUserResult struct {
		Student *HistoryUserEntity `json:"student"`
	}
	HistoryStudentUserResults struct {
		Students []*HistoryUserEntity `json:"students"`
	}

	LoginUserResult struct {
		Status string `json:"status"`
		Token  string `json:"token"`
	}
)

type (
	ActiveStudentUserEntity struct {
		ActiveStudentUserId types.IDENTIFICATION `gorm:"primaryKey"`
		Class               string               `gorm:"max:4"`
		Name                string               `gorm:"max:16"`
		Wallet              *models.ActiveWallet
		Login          *models.ActiveLogin
		Option              Options
	}
	HistoryUserEntity struct {
	}
)
