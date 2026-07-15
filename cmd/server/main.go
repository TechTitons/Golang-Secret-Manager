package main

import (
	"log"
	"net/http"

	"github.com/TechTitons/Golang-Secret-Manager/config"
	awsclient "github.com/TechTitons/Golang-Secret-Manager/pkg/aws"

	"github.com/TechTitons/Golang-Secret-Manager/internal/handler"
	"github.com/TechTitons/Golang-Secret-Manager/internal/repository"
	"github.com/TechTitons/Golang-Secret-Manager/internal/routes"
	"github.com/TechTitons/Golang-Secret-Manager/internal/service"
)

func main() {

	cfg := config.LoadConfig()

	client, err := awsclient.NewSecretManagerClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewSecretRepository(client)
	svc := service.NewSecretService(repo)
	h := handler.NewSecretHandler(svc)

	routes.RegisterRoutes(h)

	log.Println("Server Started on Port :", cfg.Port)

	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}