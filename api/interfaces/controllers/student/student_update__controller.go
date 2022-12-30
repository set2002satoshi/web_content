package student

import (
	"time"

	c "github.com/set2002satoshi/web_contents/api/interfaces/controllers"
	"github.com/set2002satoshi/web_contents/api/models"
	"github.com/set2002satoshi/web_contents/api/pkg/module/dto/request"
	"github.com/set2002satoshi/web_contents/api/pkg/module/dto/response"
	"github.com/set2002satoshi/web_contents/api/pkg/module/types"
)

type (
	UpdateActiveStudentUserResponse struct {
		response.UpdateActiveStudentUserResponse
	}
)

func (sc *StudentController) Update(ctx c.Context) {

	req := &request.ActiveStudentUpdateRequest{}
	res := &UpdateActiveStudentUserResponse{}

	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(200, "bind err")
		return
	}

	reqModel, err := UpdateFormToModel(ctx, req)
	if err != nil {
		ctx.JSON(200, "to model Err")
		return
	}
	acquired, err := sc.Interactor.Update(ctx, reqModel)
	if err != nil {
		ctx.JSON(200, err.Error())
		return
	}
	res.Result = &response.ActiveStudentUserResult{Student: sc.convertActiveStudentToDTO(acquired)}
	ctx.JSON(200, res)

}

func UpdateFormToModel(ctx c.Context, req *request.ActiveStudentUpdateRequest) (*models.ActiveStudentUser, error) {

	return models.NewActiveStudentUser(
		req.Id,
		req.Class,
		req.Name,
		nil,
		nil,
		types.REVISION(req.Revision),
		time.Time{},
		time.Time{},
	)
}
