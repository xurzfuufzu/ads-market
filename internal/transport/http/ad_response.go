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

// CreateAdResponse handles the creation of a new ad response.
// @Summary Create a new ad response
// @Description This endpoint allows the creation of a new ad response by an influencer.
// @Tags AdResponse
// @Accept  json
// @Produce  json
// @Param request body ad_response.CreateRequest true "Ad Response Creation Request"
// @Success 201 {object} fiber.Map{"message": "Ad response created successfully"}
// @Failure 400 {object} fiber.Map{"error": "Invalid request body"}
// @Failure 500 {object} fiber.Map{"error": "Internal server error"}
// @Router /ad-response/create [post]
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

// UpdateAdStatus handles the updating of an ad response's status.
// @Summary Update the status of an ad response
// @Description This endpoint allows updating the status of an existing ad response.
// @Tags AdResponse
// @Accept  json
// @Produce  json
// @Param dto body ad_response.UpdateAdStatusDTO true "Ad Response Status Update Request"
// @Success 200 {object} fiber.Map{"message": "Ad response status updated successfully"}
// @Failure 400 {object} fiber.Map{"error": "Invalid request body"}
// @Failure 500 {object} fiber.Map{"error": "Internal server error"}
// @Router /ad-response/update-status [put]
func (h *AdResponseHandler) UpdateAdStatus(c fiber.Ctx) error {
	var dto ad_response.UpdateAdStatusDTO

	if err := json.Unmarshal(c.Body(), &dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := h.adResponseService.UpdateAdResponseStatus(c.Context(), dto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Ad response status updated successfully",
	})
}
