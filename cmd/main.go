package main

import (
	"Ads-marketplace/config"
	"Ads-marketplace/internal/repository"
	"Ads-marketplace/internal/service"
	handler "Ads-marketplace/internal/transport/http"
	"Ads-marketplace/internal/transport/http/routes"
	"Ads-marketplace/pkg/client"
	"context"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"log"
)

// @title Marketplace
// @version 1.0
// @description API Server for Marketplace for influences and companies
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description JWT token
func main() {
	cfg := config.NewConfig()

	db, err := client.NewClient(context.Background(), 3, cfg.DB)
	if err != nil {
		log.Println("error to connection to postqresql", err)
	}

	if err := client.Migrate(db); err != nil {
		log.Println("migration failed", err)
	}

	defer db.Close()

	repos := repository.NewRepositories(db)
	companyService := service.NewCompanyService(repos.Company)
	companyHandler := handler.NewCompanyHandler(companyService)

	influencerService := service.NewInfluencerService(repos.Influencer)
	influencerHandler := handler.NewInfluencerHandler(influencerService)

	adResponseService := service.NewAdResponseService(repos.AdResponse)
	adResponseHandler := handler.NewAdResponseHandler(adResponseService)

	adService := service.NewAdService(repos.Ad)
	adHandler := handler.NewAdHandler(adService)

	app := fiber.New()
	app.Use(logger.New())

	routes.InitRoutes(app, companyHandler, influencerHandler, adHandler, adResponseHandler)

	log.Fatal(app.Listen(":" + "8080"))
}
