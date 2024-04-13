package service

import (
	"context"
	"golang-prisma/data/request"
	"golang-prisma/data/response"
)

type PostService interface {
	Create(ctx context.Context, requst request.PostCreateRequest)
	Update(ctx context.Context, requst request.PostUpdateRequest)
	Delete(ctx context.Context, postId string)
	FindById(ctx context.Context, postId string) response.PostResponse
	FindAll(ctx context.Context) []response.PostResponse
}
