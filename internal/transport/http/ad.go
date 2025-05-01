package handler

//
//import (
//	"Ads-marketplace/internal/domain/ad"
//	"Ads-marketplace/internal/service"
//	"github.com/gofiber/fiber/v3"
//	"net/http"
//)
//
//type AdHandler struct {
//	adService *service.AdService
//}
//
//func NewAdHandler(adService *service.AdService) *AdHandler {
//	return &AdHandler{
//		adService: adService,
//	}
//}
//
//func (h *AdHandler) CreateAd(c fiber.Ctx) error {
//	var req ad.CreateAdRequest
//	if err := c.Body(&req); err != nil {
//		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
//			"error": "Invalid input",
//		})
//	}
//
//	adEntity := &ad.Entity{
//		Title:         req.Title,
//		CompanyName:   req.CompanyName,
//		Description:   req.Description,
//		Price:         req.Price,
//		Category:      req.Category,
//		TargetCountry: req.TargetCountry,
//		Platform:      req.Platforms,
//	}
//
//	if err := h.adService.CreateAd(c.Context(), adEntity); err != nil {
//		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
//			"error": err.Error(),
//		})
//	}
//
//	return c.Status(http.StatusCreated).JSON(fiber.Map{
//		"message": "Ad created successfully",
//		"ad":      adEntity,
//	})
//}
