package student

import (

	c "github.com/set2002satoshi/web_contents/api/interfaces/controllers"
)

func (sc *StudentController) FindAll(ctx c.Context) {

	
	registeredStudent, err := sc.Interactor.FindAll(ctx)
	if err != nil {
		ctx.JSON(200, err)
		return 
	}
	ctx.JSON(200, registeredStudent)
	return
}

