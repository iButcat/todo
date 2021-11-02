package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"todo/controller"
	"todo/repository"
	"todo/service"
)

func main() {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("db failed to init")
	}

	var (
		repository = repository.NewRepository(db)
		service    = service.NewService(repository)
		controller = controller.NewController(service)
	)

	var router = gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	{
		v1 := router.Group("v1")
		{
			todoGroup := v1.Group("todo")
			{
				todoGroup.POST("/create", controller.Create)
				todoGroup.GET("/getall", controller.GetAll)
				todoGroup.GET("/getid", controller.GetByID)
				todoGroup.PUT("/update", controller.Update)
				todoGroup.DELETE("/delete", controller.Delete)
			}
		}
	}

	var errs = make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		log.Println("Starting server...")
		errs <- router.Run(":8080")
	}()

	log.Println("exit", <-errs)
}
