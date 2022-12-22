package request

type (
	ActiveStudentFindByIDRequest struct {
		Id int `json:"id"`
	}
	ActiveStudentCreateRequest struct {
		Class string `json:"class"`
		Name  string `json:"name"`
	}
	ActiveStudentUpdateRequest struct {
		Id int `json:"id"`
		Class string `json:"class"`
		Name  string `json:"name"`
	}
	ActiveStudentDeleteRequest struct {
		Id int `json:"id"`
	}
)

// type (
// 	ActiveStudentLoginRequest struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}
// )
