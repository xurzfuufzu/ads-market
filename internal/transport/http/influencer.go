package handler

import (
	"Ads-marketplace/internal/domain"
	"Ads-marketplace/internal/domain/influencer"
	"Ads-marketplace/internal/service"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"log"
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

// Register handles influencer registration.
// @Summary Register a new influencer
// @Description This endpoint allows the registration of a new influencer in the marketplace.
// @Tags Influencer
// @Accept  json
// @Produce  json
// @Param request body influencer.RegisterRequest true "Influencer Registration Request"
// @Success 200 {object} fiber.Map{"id": "influencer ID", "token": "auth token"}
// @Failure 400 {object} fiber.Map{"error": "Invalid request"}
// @Failure 500 {object} fiber.Map{"error": "Internal server error"}
// @Router /influencer/register [post]
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

// Login handles influencer login.
// @Summary Login an existing influencer
// @Description This endpoint allows an existing influencer to log in to the marketplace.
// @Tags Influencer
// @Accept  json
// @Produce  json
// @Param request body domain.LoginRequest true "Influencer Login Request"
// @Success 200 {object} fiber.Map{"id": "influencer ID", "token": "auth token"}
// @Failure 400 {object} fiber.Map{"error": "Invalid request"}
// @Failure 500 {object} fiber.Map{"error": "Internal server error"}
// @Router /influencer/login [post]
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

// GetByID handles fetching an influencer by ID.
// @Summary Get influencer by ID
// @Description This endpoint retrieves an influencer by their unique ID.
// @Tags Influencer
// @Param id path string true "Influencer ID"
// @Success 200 {object} influencer.Influencer "Influencer details"
// @Failure 400 {object} fiber.Map{"error": "id is required"}
// @Failure 500 {object} fiber.Map{"error": "Internal server error"}
// @Router /influencer/{id} [get]
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

// GetAll handles fetching all influencers.
// @Summary Get all influencers
// @Description This endpoint retrieves a list of all influencers in the marketplace.
// @Tags Influencer
// @Success 200 {array} influencer.Influencer "List of influencers"
// @Failure 500 {object} fiber.Map{"error": "Internal server error"}
// @Router /influencers [get]
func (h *InfluencerHandler) GetAll(c fiber.Ctx) error {
	influencers, err := h.influencerService.GetAllInfluencers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(influencers)
}

// GetAdsResponsesByID handles fetching all ad responses by influencer ID.
// @Summary Get all ad responses by influencer ID
// @Description This endpoint retrieves all ad responses made by a specific influencer using their ID.
// @Tags Influencer
// @Param id path string true "Influencer ID"
// @Success 200 {array} ad_response.AdResponse "List of ad responses"
// @Failure 400 {object} fiber.Map{"error": "influencer_id is required"}
// @Failure 500 {object} fiber.Map{"error": "Internal server error"}
// @Router /influencer/responses/{id} [get]
func (h *InfluencerHandler) GetAdsResponsesByID(c fiber.Ctx) error {
	influencerID := c.Params("id")
	if influencerID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "influencer_id is required",
		})
	}

	responses, err := h.influencerService.GetAdsResponsesByID(c.Context(), influencerID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if responses == nil {
		return c.JSON([]any{})
	}

	return c.JSON(responses)
}

// DeleteByID handles the deletion of an influencer by ID.
// @Summary Delete influencer by ID
// @Description This endpoint deletes an influencer by their unique ID.
// @Tags Influencer
// @Param id path string true "Influencer ID"
// @Success 200 {object} fiber.Map{}
// @Failure 400 {object} fiber.Map{"error": "ID is required"}
// @Failure 500 {object} fiber.Map{"error": "Internal server error"}
// @Router /influencer/delete/{id} [delete]
func (h *InfluencerHandler) DeleteByID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "ID is required",
		})
	}
	err := h.influencerService.DeleteByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{})
}

// Update обновляет информацию об инфлюенсере
// @Summary Обновить данные инфлюенсера
// @Description Обновляет имя, email, телефон, платформы, категорию и город инфлюенсера
// @Tags influencer
// @Accept json
// @Produce json
// @Param influencer body influencer.Entity true "Данные инфлюенсера для обновления"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string{"error":"Invalid request"}
// @Failure 500 {object} map[string]string{"error":"internal server error"}
// @Router /influencer/update [put]
func (h *InfluencerHandler) Update(c fiber.Ctx) error {
	var influencer influencer.Entity
	if err := json.Unmarshal(c.Body(), &influencer); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	log.Println("input", influencer)

	err := h.influencerService.UpdateByID(c.Context(), &influencer)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{})
}
