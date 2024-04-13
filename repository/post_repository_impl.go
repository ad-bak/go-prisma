package repository

import (
	"context"
	"errors"
	"fmt"
	"golang-prisma/helper"
	"golang-prisma/model"
	"golang-prisma/prisma/db"
)

type PostRepositoryImpl struct {
	DB *db.PrismaClient
}

func NewPostRepository(Db *db.PrismaClient) PostRepository {
	return &PostRepositoryImpl{DB: Db}
}

func (p *PostRepositoryImpl) Delete(ctx context.Context, postId string) {
	result, err := p.DB.Post.FindUnique(db.Post.ID.Equals(postId)).Delete().Exec(ctx)
	helper.ErrorPanic(err)
	fmt.Println("Rows affected:", result)
}

func (p *PostRepositoryImpl) FindAll(ctx context.Context) []model.Post {
	allPost, err := p.DB.Post.FindMany().Exec(ctx)
	helper.ErrorPanic(err)

	var posts []model.Post

	for _, post := range allPost {
		published, _ := post.Published()
		description, _ := post.Description()

		postData := model.Post{
			Id:          post.ID,
			Title:       post.Title,
			Published:   published,
			Description: description,
		}
		posts = append(posts, postData)
	}
	return posts
}

func (p *PostRepositoryImpl) FindById(ctx context.Context, postId string) (model.Post, error) {
	post, err := p.DB.Post.FindFirst(db.Post.ID.Equals(postId)).Exec(ctx)
	helper.ErrorPanic(err)

	published, _ := post.Published()
	description, _ := post.Description()
	postData := model.Post{
		Id:          post.ID,
		Title:       post.Title,
		Published:   published,
		Description: description,
	}

	if post != nil {
		return postData, nil
	} else {
		return postData, errors.New("Post not found")
	}
}

func (p *PostRepositoryImpl) Save(ctx context.Context, post model.Post) {
	result, err := p.DB.Post.CreateOne(
		db.Post.Title.Set(post.Title),
		db.Post.Published.Set(post.Published),
		db.Post.Description.Set(post.Description),
	).Exec(ctx)
	helper.ErrorPanic(err)
	fmt.Println("Post created with ID:", result.ID)
}

func (p *PostRepositoryImpl) Update(ctx context.Context, post model.Post) {
	result, err := p.DB.Post.FindMany(db.Post.ID.Equals(post.Id)).Update(
		db.Post.Title.Set(post.Title),
		db.Post.Published.Set(post.Published),
		db.Post.Description.Set(post.Description),
	).Exec(ctx)
	helper.ErrorPanic(err)
	fmt.Println("Post created with ID:", result)
}
