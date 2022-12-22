package student

import (
	"time"

	c "github.com/set2002satoshi/web_contents/api/interfaces/controllers"
	"github.com/set2002satoshi/web_contents/api/models"
	"github.com/set2002satoshi/web_contents/api/pkg/module/dto/request"
	"github.com/set2002satoshi/web_contents/api/pkg/module/types"
)

func (sc *StudentController) Create(ctx c.Context) {

	req := &request.ActiveStudentCreateRequest{}

	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(200, "bind err")
		return 
	}

	reqModel, err := CreateFormToModel(ctx, req)
	if err != nil {
		ctx.JSON(200, "to model Err")
		return 
	}
	registeredStudent, err := sc.Interactor.Register(ctx, reqModel)
	if err != nil {
		ctx.JSON(200, err)
		return 
	}
	ctx.JSON(200, registeredStudent)
	return
}

func CreateFormToModel(ctx c.Context, req *request.ActiveStudentCreateRequest) (*models.ActiveStudentUser, error) {

	coin := 0
	auth, err := models.NewActiveStudentAuth(
		types.INITIAL_ID,
		req.Email,
		req.Password,
	)
	if err != nil {
		return nil, err
	}
	wallet, err := models.NewActiveWallet(
		types.INITIAL_ID,
		types.INITIAL_ID,
		uint(coin),
	)
	if err != nil {
		return nil, err
	}
	return models.NewActiveStudentUser(
		types.INITIAL_ID,
		req.Class,
		req.Name,
		auth,
		wallet,
		time.Time{},
		time.Time{},
	)
}