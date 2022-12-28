package infrastructure

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	sc "github.com/set2002satoshi/web_contents/api/interfaces/controllers/student"
)

type Routing struct {
	DB   *DB
	Gin  *gin.Engine
	Port string
}

func NewRouting(db *DB) *Routing {
	r := &Routing{
		DB:   db,
		Gin:  gin.Default(),
		Port: ":8080",
	}
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {

	studentController := sc.NewStudentController(r.DB)


	r.Gin.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	})


	student := r.Gin.Group("/api")
	{
		student.POST("/find/byId", func(c *gin.Context) {studentController.FindById(c)})
		student.POST("/find/all", func(c *gin.Context) {studentController.FindAll(c)})
		student.POST("/create",func(c *gin.Context) {studentController.Create(c)})
		student.POST("/delete", func(c *gin.Context) {studentController.Delete(c)})
	}


}

func (r *Routing) Run() {
	err := r.Gin.Run(r.Port)
	if err != nil {
		log.Fatalln("err:", err)
	}
}
