package router

import (
	"fmt"
	"golang-prisma/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(postController *controller.PostController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Print("Welcome to Golang Prisma")
	})

	router.GET("/posts", postController.FindAll)
	router.GET("/posts/:postId", postController.FindById)
	router.POST("/posts", postController.Create)
	router.PATCH("/posts/:postId", postController.Update)
	router.DELETE("/posts/:postId", postController.Delete)

	return router

}
