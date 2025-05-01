package service

import (
	"Ads-marketplace/internal/domain"
	"Ads-marketplace/internal/domain/influencer"
	"Ads-marketplace/internal/repository"
	"Ads-marketplace/pkg/hasher"
	"Ads-marketplace/pkg/token"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type InfluencerService struct {
	influencerRepo repository.Influencer
}

func NewInfluencerService(influencerRepo repository.Influencer) *InfluencerService {
	return &InfluencerService{
		influencerRepo: influencerRepo,
	}
}

func (s *InfluencerService) Register(ctx context.Context, input influencer.RegisterRequest) (string, string, error) {
	existingInfluencer, err := s.influencerRepo.GetByEmail(ctx, input.Email)
	if err == nil && existingInfluencer != nil {
		return "", "", influencer.ErrorEmailConflict
	}

	hashedPassword, err := hasher.Hash(input.Password)
	if err != nil {
		return "", "", err
	}

	newInfluencer := &influencer.Entity{
		ID:          uuid.New(),
		Name:        input.Name,
		Email:       input.Email,
		Password:    hashedPassword,
		Phone:       input.PhoneNumber,
		AccountType: "Influencer",
	}

	influencerID, err := s.influencerRepo.Create(ctx, newInfluencer)
	if err != nil {
		return "", "", err
	}

	token, err := token.GenerateToken(influencerID)
	if err != nil {
		return "", "", err
	}

	return influencerID, token, nil
}

func (s *InfluencerService) Login(ctx context.Context, input domain.LoginRequest) (string, string, error) {
	existingInfluencer, err := s.influencerRepo.GetByEmail(ctx, input.Email)
	if err != nil || existingInfluencer == nil {
		return "", "", errors.New("influencer not found")
	}

	err = hasher.Compare(existingInfluencer.Password, input.Password)
	if err != nil {
		return "", "", errors.New("invalid password")
	}

	token, err := token.GenerateToken(existingInfluencer.ID.String())
	if err != nil {
		return "", "", err
	}

	return existingInfluencer.ID.String(), token, nil
}

func (s *InfluencerService) GetByID(ctx context.Context, id string) (*influencer.Entity, error) {
	return s.influencerRepo.GetByID(ctx, id)
}

func (s *InfluencerService) GetAllInfluencers(ctx context.Context) ([]*influencer.Entity, error) {
	influencers, err := s.influencerRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all influencers: %v", err)
	}
	return influencers, nil
}
