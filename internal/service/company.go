package service

import (
	"Ads-marketplace/internal/domain"
	"Ads-marketplace/internal/domain/ad"
	"Ads-marketplace/internal/domain/company"
	"Ads-marketplace/internal/domain/influencer"
	"Ads-marketplace/internal/repository"
	"Ads-marketplace/pkg/hasher"
	"Ads-marketplace/pkg/token"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type CompanyService struct {
	companyRepo repository.Company
}

func NewCompanyService(companyRepo repository.Company) *CompanyService {
	return &CompanyService{
		companyRepo: companyRepo,
	}
}

func (s *CompanyService) Register(ctx context.Context, input company.RegisterRequest) (string, string, error) {
	existingCompany, err := s.companyRepo.GetByEmail(ctx, input.Email)
	if err == nil && existingCompany != nil {
		return "", "", company.ErrorEmailConflict
	}

	hashedPassword, err := hasher.Hash(input.Password)

	company := &company.Entity{
		ID:          uuid.New(),
		Name:        input.Name,
		Email:       input.Email,
		Password:    hashedPassword,
		Phone:       input.PhoneNumber,
		AccountType: "Company",
	}

	id, err := s.companyRepo.Create(ctx, company)
	if err != nil {
		return "", "", err
	}

	token, err := token.GenerateToken(company.ID.String())
	if err != nil {
		return "", "", err
	}

	return id, token, nil
}

func (s *CompanyService) Login(ctx context.Context, input domain.LoginRequest) (string, string, error) {
	existingCompany, err := s.companyRepo.GetByEmail(ctx, input.Email)
	if err != nil {
		return "", "", errors.New("company not found")
	}

	if err := hasher.Compare(existingCompany.Password, input.Password); err != nil {
		return "", "", errors.New("invalid credentials")
	}

	token, err := token.GenerateToken(existingCompany.ID.String())
	if err != nil {
		return "", "", err
	}

	return existingCompany.ID.String(), token, nil
}

func (s *CompanyService) GetByID(ctx context.Context, id string) (*company.Entity, error) {
	return s.companyRepo.GetByID(ctx, id)
}

func (s *CompanyService) GetAdsByCompanyName(ctx context.Context, companyName string) ([]*ad.Entity, error) {
	return s.companyRepo.GetCompanyAds(ctx, companyName)
}

func (s *CompanyService) DeleteByID(ctx context.Context, id string) error {
	err := s.companyRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete ad: %v", err)
	}

	return nil
}

func (s *CompanyService) GetAdResponses(ctx context.Context, companyID uuid.UUID) ([]*influencer.InfluencerDTO, error) {
	ads, err := s.companyRepo.GetInfluencersForAd(ctx, companyID)
	if err != nil {
		return nil, fmt.Errorf("failed to get ad responses: %v", err)
	}

	return ads, nil
}

func (s *CompanyService) UpdateByID(ctx context.Context, entity *company.Entity) error {
	err := s.companyRepo.Update(ctx, entity)
	if err != nil {
		return fmt.Errorf("failed to update company: %v", err)
	}

	return nil
}
