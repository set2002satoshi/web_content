package student

import (
	DBstudent "github.com/set2002satoshi/web_contents/api/interfaces/database/student"
	database "github.com/set2002satoshi/web_contents/api/interfaces/database"
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
