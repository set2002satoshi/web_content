package student

import (
	c "github.com/set2002satoshi/web_contents/api/interfaces/controllers"
	"github.com/set2002satoshi/web_contents/api/pkg/module/dto/response"
)

type (
	FindAllActiveStudentUserResponse struct {
		response.FindAllActiveStudentUserResponse
	}
)

func (sc *StudentController) FindAll(ctx c.Context) {

	res := &FindAllActiveStudentUserResponse{}
	
	acquired, err := sc.Interactor.FindAll(ctx)
	if err != nil {
		ctx.JSON(200, err)
		return 
	}

	res.Result = &response.ActiveStudentUserResults{Students: sc.convertActiveStudentToDTOs(acquired)}
	ctx.JSON(200, res)
}

