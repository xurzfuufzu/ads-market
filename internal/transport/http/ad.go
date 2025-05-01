package handler

import (
	"Ads-marketplace/internal/domain/ad"
	"Ads-marketplace/internal/service"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"net/http"
)

type AdHandler struct {
	adService *service.AdService
}

func NewAdHandler(adService *service.AdService) *AdHandler {
	return &AdHandler{
		adService: adService,
	}
}

func (h *AdHandler) Create(c fiber.Ctx) error {
	var input *ad.CreateRequest
	if err := json.Unmarshal(c.Body(), &input); err != nil {
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if err := h.adService.CreateAd(c.Context(), input); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Ad created successfully",
	})
}

func (h *AdHandler) GetAllAds(c fiber.Ctx) error {
	ads, err := h.adService.GetAllAds(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(ads)
}
