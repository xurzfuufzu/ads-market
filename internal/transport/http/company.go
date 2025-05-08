package handler

import (
	"Ads-marketplace/internal/domain"
	"Ads-marketplace/internal/domain/company"
	"Ads-marketplace/internal/service"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
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

// Register handles company registration.
// @Summary Register a new company
// @Description This endpoint allows the registration of a new company in the marketplace.
// @Tags Company
// @Accept  json
// @Produce  json
// @Param request body company.RegisterRequest true "Company Registration Request"
// @Success 200 {object} {"id": "company ID", "token": "auth token"}
// @Failure 400 {object} {"error": "Invalid request"}
// @Failure 500 {object} {"error": "Internal server error"}
// @Router /company/register [post]
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

// Login handles company login.
// @Summary Login an existing company
// @Description This endpoint allows an existing company to log in to the marketplace.
// @Tags Company
// @Accept  json
// @Produce  json
// @Param request body domain.LoginRequest true "Company Login Request"
// @Success 200 {object} fiber.Map{"id": "company ID", "token": "auth token"}
// @Failure 400 {object} fiber.Map{"error": "Invalid request"}
// @Failure 401 {object} fiber.Map{"error": "Unauthorized"}
// @Failure 500 {object} fiber.Map{"error": "Internal server error"}
// @Router /company/login [post]
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

// GetByID handles fetching a company by its ID.
// @Summary Get company by ID
// @Description This endpoint retrieves a company by its unique ID.
// @Tags Company
// @Param id path string true "Company ID"
// @Success 200 {object} company.Company "Company details"
// @Failure 400 {object} fiber.Map{"error": "id is required"}
// @Failure 500 {object} fiber.Map{"error": "Internal server error"}
// @Router /company/{id} [get]
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

// GetAdsByCompanyName handles fetching all ads by company name.
// @Summary Get all ads by company name
// @Description This endpoint retrieves all advertisements posted by a specific company using the company's name.
// @Tags Company
// @Param company_name path string true "Company name"
// @Success 200 {array} ad.Ad "List of advertisements"
// @Failure 400 {object} fiber.Map{"error": "company_name is required"}
// @Failure 500 {object} fiber.Map{"error": "Internal server error"}
// @Router /company/ads/{company_name} [get]
func (h *CompanyHandler) GetAdsByCompanyName(c fiber.Ctx) error {
	companyName := c.Params("company_name")
	if companyName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "company_name is required",
		})
	}

	ads, err := h.companyService.GetAdsByCompanyName(c.Context(), companyName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(ads)
}

// DeleteByID handles the deletion of a company by ID.
// @Summary Delete company by ID
// @Description This endpoint deletes a company by its unique ID.
// @Tags Company
// @Param id path string true "Company ID"
// @Success 200 {object} fiber.Map{}
// @Failure 400 {object} fiber.Map{"error": "ID is required"}
// @Failure 500 {object} fiber.Map{"error": "Internal server error"}
// @Router /company/delete/{id} [delete]
func (h *CompanyHandler) DeleteByID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "ID is required",
		})
	}
	err := h.companyService.DeleteByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{})
}

// GetResponsesByAdID handles fetching all responses to a specific ad by its ID.
// @Summary Get all responses to an ad by ad ID
// @Description This endpoint retrieves all responses to a particular ad using its ID.
// @Tags Company
// @Param ad_id path string true "Ad ID"
// @Success 200 {array} ad_response.AdResponse "List of ad responses"
// @Failure 400 {object} fiber.Map{"error": "ad_id is required"}
// @Failure 500 {object} fiber.Map{"error": "Internal server error"}
// @Router /company/responses/{ad_id} [get]
func (h *CompanyHandler) GetResponsesByAdID(c fiber.Ctx) error {
	adID := c.Params("ad_id")
	if adID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "ad_id is required",
		})
	}

	fmt.Println("adID", adID)

	id, err := uuid.Parse(adID)
	responses, err := h.companyService.GetAdResponses(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(responses)
}

// Update обновляет информацию о компании
// @Summary Обновить данные компании
// @Description Обновляет имя, email, телефон, описание и город компании
// @Tags company
// @Accept json
// @Produce json
// @Param company body company.Entity true "Данные компании для обновления"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string{"error":"Invalid request"}
// @Failure 500 {object} map[string]string{"error":"internal server error"}
// @Router /company/update [put]
func (h *CompanyHandler) Update(c fiber.Ctx) error {
	var company company.Entity
	if err := json.Unmarshal(c.Body(), &company); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}
	err := h.companyService.UpdateByID(c.Context(), &company)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{})
}
