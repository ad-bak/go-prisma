package controller

import (
	"golang-prisma/data/request"
	"golang-prisma/data/response"
	"golang-prisma/helper"
	"golang-prisma/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PostController struct {
	PostService service.PostService
}

func NewPostController(postService service.PostService) *PostController {
	return &PostController{PostService: postService}
}

func (controller *PostController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	postCreateRequest := request.PostCreateRequest{}
	helper.ReadRequestBody(requests, &postCreateRequest)

	controller.PostService.Create(requests.Context(), postCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteResponseBody(writer, webResponse)

}

func (controller *PostController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	postUpdateRequest := request.PostUpdateRequest{}
	helper.ReadRequestBody(requests, &postUpdateRequest)

	postId := params.ByName("postId")
	postUpdateRequest.Id = postId

	controller.PostService.Update(requests.Context(), postUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteResponseBody(writer, webResponse)

}

func (controller *PostController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {

	result := controller.PostService.FindAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}
	helper.WriteResponseBody(writer, webResponse)

}

func (controller *PostController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	postId := params.ByName("postId")
	result := controller.PostService.FindById(requests.Context(), postId)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}
	helper.WriteResponseBody(writer, webResponse)

}

func (controller *PostController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	postId := params.ByName("postId")
	controller.PostService.Delete(requests.Context(), postId)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteResponseBody(writer, webResponse)

}
