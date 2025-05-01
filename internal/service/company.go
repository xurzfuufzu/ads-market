package service

import (
	"Ads-marketplace/internal/domain"
	"Ads-marketplace/internal/domain/company"
	"Ads-marketplace/internal/repository"
	"Ads-marketplace/pkg/hasher"
	"Ads-marketplace/pkg/token"
	"context"
	"errors"
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

func (s *CompanyService) Register(ctx context.Context, input company.RegisterRequest) (string, error) {
	existingCompany, err := s.companyRepo.GetByEmail(ctx, input.Email)
	if err == nil && existingCompany != nil {
		return "", company.ErrorEmailConflict
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

	err = s.companyRepo.Create(ctx, company)
	if err != nil {
		return "", err
	}

	token, err := token.GenerateToken(company.ID.String())
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *CompanyService) Login(ctx context.Context, input domain.LoginRequest) (string, error) {
	existingCompany, err := s.companyRepo.GetByEmail(ctx, input.Email)
	if err != nil {
		return "", errors.New("company not found")
	}

	if err := hasher.Compare(existingCompany.Password, input.Password); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := token.GenerateToken(existingCompany.ID.String())
	if err != nil {
		return "", err
	}

	return token, nil
}
