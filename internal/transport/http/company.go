package handler

import (
	"Ads-marketplace/internal/domain"
	"Ads-marketplace/internal/domain/company"
	"Ads-marketplace/internal/service"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"net/http"
)

type CompanyHandler struct {
	companyService *service.CompanyService
}

func NewCompanyHandler(companyService *service.CompanyService) *CompanyHandler {
	return &CompanyHandler{
		companyService: companyService,
	}
}

func (h *CompanyHandler) Register(c fiber.Ctx) error {
	var input company.RegisterRequest
	if err := json.Unmarshal(c.Body(), &input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	id, token, err := h.companyService.Register(c.Context(), input)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"id":    id,
		"token": token,
	})
}

func (h *CompanyHandler) Login(c fiber.Ctx) error {
	var input domain.LoginRequest
	if err := json.Unmarshal(c.Body(), &input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	id, token, err := h.companyService.Login(c.Context(), input)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"id":    id,
		"token": token,
	})
}

func (h *CompanyHandler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	company, err := h.companyService.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(company)
}
