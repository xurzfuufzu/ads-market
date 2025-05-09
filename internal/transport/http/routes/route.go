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
	company.Get("/:company_name/ads", AuthMiddleware, companyHandler.GetAdsByCompanyName)
	company.Delete("/:id", AuthMiddleware, companyHandler.DeleteByID)
	company.Get("/:ad_id/influencers", AuthMiddleware, companyHandler.GetResponsesByAdID)
	company.Put("/update", AuthMiddleware, companyHandler.Update)

	influencer := app.Group("/influencer")
	influencer.Post("/register", influencerHandler.Register)
	influencer.Post("/login", influencerHandler.Login)
	influencer.Get("/all", AuthMiddleware, influencerHandler.GetAll)
	influencer.Get("/:id", AuthMiddleware, influencerHandler.GetByID)
	influencer.Get("/:id/responses", AuthMiddleware, influencerHandler.GetAdsResponsesByID)
	influencer.Delete("/:id", AuthMiddleware, influencerHandler.DeleteByID)
	influencer.Put("/update", AuthMiddleware, influencerHandler.Update)

	ad := app.Group("/ad")
	ad.Post("/create", AuthMiddleware, adHandler.Create)
	ad.Get("/all", AuthMiddleware, adHandler.GetAllAds)
	ad.Delete("/:id", AuthMiddleware, adHandler.DeleteByID)
	ad.Put("/update", AuthMiddleware, adHandler.UpdateByID)
	ad.Get("/:id", AuthMiddleware, adHandler.GetByID)

	adResponse := app.Group("/ad-response")
	adResponse.Post("/create", AuthMiddleware, adResponseHandler.CreateAdResponse)
	adResponse.Put("/update-status", AuthMiddleware, adResponseHandler.UpdateAdStatus)
}
