package handler

import (
	"Ads-marketplace/internal/domain/ad"
	"Ads-marketplace/internal/service"
	"Ads-marketplace/pkg/utils/response"
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
// @Success 200 {object} response.Response{Message="Ad created successfully"}
// @Failure 400 {object} response.Response{Error="Invalid request"}
// @Failure 500 {object} response.Response{Error="Internal server error"}
// @Router /ad/create [post]
func (h *AdHandler) Create(c fiber.Ctx) error {
	var input *ad.CreateRequest
	if err := json.Unmarshal(c.Body(), &input); err != nil {
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(response.ClientResponse(http.StatusBadRequest, "Invalid request", nil, err.Error()))
	}

	if err := h.adService.CreateAd(c.Context(), input); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ClientResponse(http.StatusInternalServerError, "Internal server error", nil, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(response.ClientResponse(http.StatusOK, "Ad created successfully", nil, nil))
}

// GetAllAds handles fetching all advertisements.
// @Summary Get all advertisements
// @Description This endpoint returns a list of all advertisements in the marketplace.
// @Tags Ad
// @Produce  json
// @Success 200 {array} ad.Ad "List of advertisements"
// @Failure 500 {object} response.Response{Error="Internal server error"}
// @Router /ad/all [get]
func (h *AdHandler) GetAllAds(c fiber.Ctx) error {
	ads, err := h.adService.GetAllAds(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ClientResponse(http.StatusInternalServerError, "Internal server error", nil, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(response.ClientResponse(http.StatusOK, "Success", ads, nil))
}

// DeleteByID handles the deletion of an advertisement by ID.
// @Summary Delete an advertisement by ID
// @Description This endpoint deletes an advertisement from the marketplace by its ID.
// @Tags Ad
// @Param id path string true "Ad ID"
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{Error="ID is required"}
// @Failure 500 {object} response.Response{Error="Internal server error"}
// @Router /ad/delete/{id} [delete]
func (h *AdHandler) DeleteByID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(response.ClientResponse(http.StatusBadRequest, "ID is required", nil, nil))
	}
	err := h.adService.DeleteAdByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ClientResponse(http.StatusInternalServerError, "Internal server error", nil, err.Error()))
	}
	return c.Status(http.StatusOK).JSON(response.ClientResponse(http.StatusOK, "Success", nil, nil))
}
