package service

import (
	"Ads-marketplace/internal/domain/ad"
	"Ads-marketplace/internal/repository"
	"context"
)

type AdService struct {
	adRepo repository.Ad
}

func NewAdService(adRepo repository.Ad) *AdService {
	return &AdService{
		adRepo: adRepo,
	}
}

func (s *AdService) CreateAd(ctx context.Context, ad *ad.Entity) error {
	ad.Status = "available"

	err := s.adRepo.Create(ctx, ad)
	if err != nil {
		return err
	}

	return nil
}
