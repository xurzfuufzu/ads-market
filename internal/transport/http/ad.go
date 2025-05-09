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

// Create handles the creation of a new advertisement.
// @Summary Create a new advertisement
// @Description This endpoint allows the creation of a new advertisement in the marketplace.
// @Tags Ad
// @Accept  json
// @Produce  json
// @Param request body ad.CreateRequest true "Ad Creation Request"
// @Success 200 {object} response.Message
// @Failure 400 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /ad/create [post]
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

// GetAllAds handles fetching all advertisements.
// @Summary Get all advertisements
// @Description This endpoint returns a list of all advertisements in the marketplace.
// @Tags Ad
// @Produce  json
// @Success 200 {array} ad.Ad "List of advertisements"
// @Failure 500 {object} fiber.Map{"error": "Internal server error"}
// @Router /ad/all [get]
func (h *AdHandler) GetAllAds(c fiber.Ctx) error {
	ads, err := h.adService.GetAllAds(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(ads)
}

// DeleteByID handles the deletion of an advertisement by ID.
// @Summary Delete an advertisement by ID
// @Description This endpoint deletes an advertisement from the marketplace by its ID.
// @Tags Ad
// @Param id path string true "Ad ID"
// @Success 200 {object} fiber.Map{}
// @Failure 400 {object} fiber.Map{"error": "ID is required"}
// @Failure 500 {object} fiber.Map{"error": "Internal server error"}
// @Router /ad/delete/{id} [delete]
func (h *AdHandler) DeleteByID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "ID is required",
		})
	}
	err := h.adService.DeleteAdByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{})
}

func (h *AdHandler) UpdateByID(c fiber.Ctx) error {
	var ad ad.Entity
	if err := json.Unmarshal(c.Body(), &ad); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	err := h.adService.Update(c.Context(), &ad)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{})
}
