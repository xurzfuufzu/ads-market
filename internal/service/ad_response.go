package service

import (
	"Ads-marketplace/internal/domain/ad_response"
	"Ads-marketplace/internal/repository"
	"context"
	"fmt"
)

type AdResponseService struct {
	adResponseRepo repository.AdResponse
}

func NewAdResponseService(AdResponse repository.AdResponse) *AdResponseService {
	return &AdResponseService{
		adResponseRepo: AdResponse,
	}
}

func (s *AdResponseService) CreateAdResponse(ctx context.Context, adResponse *ad_response.CreateRequest) error {
	err := s.adResponseRepo.Create(ctx, adResponse)
	if err != nil {
		return fmt.Errorf("failed to create ad response: %v", err)
	}
	return nil
}
