package student

import (
	c "github.com/set2002satoshi/web_contents/api/interfaces/controllers"
	"github.com/set2002satoshi/web_contents/api/pkg/module/dto/request"
	"github.com/set2002satoshi/web_contents/api/pkg/module/dto/response"
)

type (
	FindByIdActiveStudentUserResponse struct {
		response.FindByIDActiveStudentUserResponse
	}
)

func (sc *StudentController) FindById(ctx c.Context) {

	req := &request.ActiveStudentFindByIDRequest{}
	res := FindByIdActiveStudentUserResponse{}

	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(200, "bind err")
		return
	}
	acquired, err := sc.Interactor.FindById(ctx, req.Id)
	if err != nil {
		ctx.JSON(200, err)
		return
	}

	res.Result = &response.ActiveStudentUserResult{Student: sc.convertActiveStudentToDTO(acquired)}
	ctx.JSON(200, res)

}
