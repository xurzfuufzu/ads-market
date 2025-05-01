package handler

import (
	"Ads-marketplace/internal/domain"
	"Ads-marketplace/internal/domain/influencer"
	"Ads-marketplace/internal/service"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"net/http"
)

type InfluencerHandler struct {
	influencerService *service.InfluencerService
}

func NewInfluencerHandler(influencerService *service.InfluencerService) *InfluencerHandler {
	return &InfluencerHandler{
		influencerService: influencerService,
	}
}

func (h *InfluencerHandler) Register(c fiber.Ctx) error {
	var input influencer.RegisterRequest
	if err := json.Unmarshal(c.Body(), &input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	influencerID, token, err := h.influencerService.Register(c.Context(), input)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"id":    influencerID,
		"token": token,
	})
}

func (h *InfluencerHandler) Login(c fiber.Ctx) error {
	var input domain.LoginRequest
	if err := json.Unmarshal(c.Body(), &input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	influencerID, token, err := h.influencerService.Login(c.Context(), input)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"id":    influencerID,
		"token": token,
	})
}

func (h *InfluencerHandler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	influencer, err := h.influencerService.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(influencer)
}

func (h *InfluencerHandler) GetAll(c fiber.Ctx) error {
	influencers, err := h.influencerService.GetAllInfluencers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(influencers)
}
