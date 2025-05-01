package routes

import (
	"Ads-marketplace/internal/transport/http"
	"github.com/gofiber/fiber/v3"
)

func InitRoutes(
	app *fiber.App,
	companyHandler *handler.CompanyHandler,
	influencerHandler *handler.InfluencerHandler,
) {
	company := app.Group("/company")

	company.Post("/register", companyHandler.Register)
	company.Post("/login", companyHandler.Login)

	influencer := app.Group("/influencer")
	influencer.Post("/register", influencerHandler.Register)
	influencer.Post("/login", influencerHandler.Login)

}
