package routes

import (
	"Ads-marketplace/internal/transport/http"
	"github.com/gofiber/fiber/v3"
)

func InitRoutes(
	app *fiber.App,
	companyHandler *handler.CompanyHandler,
	influencerHandler *handler.InfluencerHandler,
	adHandler *handler.AdHandler,
	adResponseHandler *handler.AdResponseHandler,
) {
	company := app.Group("/company")
	company.Post("/register", companyHandler.Register)
	company.Post("/login", companyHandler.Login)
	company.Get("/:id", AuthMiddleware, companyHandler.GetByID)

	influencer := app.Group("/influencer")
	influencer.Post("/register", influencerHandler.Register)
	influencer.Post("/login", influencerHandler.Login)
	influencer.Get("/all", AuthMiddleware, influencerHandler.GetAll)
	influencer.Get("/:id", AuthMiddleware, influencerHandler.GetByID)

	ad := app.Group("/ad")
	ad.Post("/create", AuthMiddleware, adHandler.Create)
	ad.Get("/all", AuthMiddleware, adHandler.GetAllAds)

	adResponse := app.Group("/ad-response")
	adResponse.Post("/create", AuthMiddleware, adResponseHandler.CreateAdResponse)
}
