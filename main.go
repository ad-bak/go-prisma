package main

import (
	"fmt"
	"golang-prisma/config"
	"golang-prisma/controller"
	"golang-prisma/helper"
	"golang-prisma/repository"
	"golang-prisma/router"
	"golang-prisma/service"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	fmt.Print("Start server!" + os.Getenv("PORT"+"\n"))

	db, err := config.ConnectDB()
	helper.ErrorPanic(err)

	defer db.Prisma.Disconnect()

	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostServiceImpl(postRepository)
	postContoller := controller.NewPostController(postService)
	routes := router.NewRouter(postContoller)

	server := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        routes,
	}

	server_err := server.ListenAndServe()
	if server_err != nil {
		panic(server_err)
	}

}
