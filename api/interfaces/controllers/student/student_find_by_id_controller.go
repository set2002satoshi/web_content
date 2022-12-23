package student

import (

	c "github.com/set2002satoshi/web_contents/api/interfaces/controllers"
	"github.com/set2002satoshi/web_contents/api/pkg/module/dto/request"
)

func (sc *StudentController) FindById(ctx c.Context) {

	req := &request.ActiveStudentFindByIDRequest{}

	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(200, "bind err")
		return 
	}
	registeredStudent, err := sc.Interactor.FindById(ctx, req.Id)
	if err != nil {
		ctx.JSON(200, err)
		return 
	}
	ctx.JSON(200, registeredStudent)
	return
}

