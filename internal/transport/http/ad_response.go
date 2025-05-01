package handler

import (
	"Ads-marketplace/internal/domain/ad_response"
	"Ads-marketplace/internal/service"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"net/http"
)

type AdResponseHandler struct {
	adResponseService *service.AdResponseService
}

func NewAdResponseHandler(adResponseService *service.AdResponseService) *AdResponseHandler {
	return &AdResponseHandler{
		adResponseService: adResponseService,
	}
}

func (h *AdResponseHandler) CreateAdResponse(c fiber.Ctx) error {
	var request ad_response.CreateRequest

	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := h.adResponseService.CreateAdResponse(c.Context(), &request)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Ad response created successfully",
	})
}
