package main

import (
	"log"
	"net/http"

	"github.com/dieg0code/products-microservice/src/controllers"
	"github.com/dieg0code/products-microservice/src/db"
	"github.com/dieg0code/products-microservice/src/models"
	"github.com/dieg0code/products-microservice/src/repository"
	"github.com/dieg0code/products-microservice/src/router"
	"github.com/dieg0code/products-microservice/src/services"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func main() {
	db := db.DatabaseConnection()
	err := db.AutoMigrate(&models.Product{})
	if err != nil {
		logrus.Fatalf("Failed to migrate database: %v", err)
		panic("Failed to migrate database")
	}

	repo := repository.NewPorductRespositoryImpl(db)

	service := services.NewProductServiceImpl(repo)

	validator := validator.New()

	controller := controllers.NewProductControllerImpl(service, validator)

	r := router.NewRouter(controller)

	ginRouter := r.InitRoutes()

	server := &http.Server{
		Addr:    ":8080",
		Handler: ginRouter,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
