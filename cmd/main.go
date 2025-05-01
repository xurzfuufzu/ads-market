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

func main() {
	cfg := config.NewConfig()

	db, err := client.NewClient(context.Background(), 3, cfg.DB)
	if err != nil {
		log.Println("error to connection to postqresql", err)
	}

	defer db.Close()

	repos := repository.NewRepositories(db)
	companyService := service.NewCompanyService(repos.Company)
	companyHandler := handler.NewCompanyHandler(companyService)

	influencerService := service.NewInfluencerService(repos.Influencer)
	influencerHandler := handler.NewInfluencerHandler(influencerService)

	app := fiber.New()
	app.Use(logger.New())

	routes.InitRoutes(app, companyHandler, influencerHandler)

	log.Fatal(app.Listen(":" + cfg.Server.Port))
}
