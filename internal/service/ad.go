package service

import (
	"Ads-marketplace/internal/domain/ad"
	"Ads-marketplace/internal/repository"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type AdService struct {
	adRepo repository.Ad
}

func NewAdService(adRepo repository.Ad) *AdService {
	return &AdService{
		adRepo: adRepo,
	}
}

func (s *AdService) CreateAd(ctx context.Context, adRequest *ad.CreateRequest) error {
	ad := &ad.Entity{
		ID:          uuid.New().String(), // Генерируем ID для объявления
		Title:       adRequest.Title,
		CompanyName: adRequest.CompanyName,
		Description: adRequest.Description,
		PriceFrom:   adRequest.PriceFrom,
		PriceTo:     adRequest.PriceTo,
		Status:      "active",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Platforms:   adRequest.Platforms,
		Category:    adRequest.Category,
		City:        adRequest.City,
	}

	err := s.adRepo.Create(ctx, ad)
	if err != nil {
		return fmt.Errorf("failed to create ad: %v", err)
	}

	return nil
}

func (s *AdService) GetAllAds(ctx context.Context) ([]*ad.Entity, error) {
	ads, err := s.adRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get ads: %v", err)
	}

	return ads, nil
}

func (s *AdService) DeleteAdByID(ctx context.Context, id string) error {
	err := s.adRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete ad: %v", err)
	}

	return nil
}

func (s *AdService) Update(ctx context.Context, ad *ad.Entity) error {
	if ad.ID == "" {
		return errors.New("ad ID is required")
	}

	return s.adRepo.Update(ctx, ad)
}
