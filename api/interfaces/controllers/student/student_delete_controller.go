package student

import (
	c "github.com/set2002satoshi/web_contents/api/interfaces/controllers"
	"github.com/set2002satoshi/web_contents/api/pkg/module/dto/request"
)

func (sc *StudentController) Delete(ctx c.Context) {

	req := &request.ActiveStudentDeleteRequest{}

	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(200, "bind err")
		return
	}
	err := sc.Interactor.DeleteById(req.Id)
	if err != nil {
		ctx.JSON(200, err)
		return
	}

	ctx.JSON(200, "削除完了")
}