package repository

import (
	"Ads-marketplace/internal/domain/ad"
	"Ads-marketplace/internal/domain/ad_response"
	"Ads-marketplace/internal/domain/company"
	"Ads-marketplace/internal/domain/influencer"
	"Ads-marketplace/internal/repository/postgresql"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Ad interface {
	Create(ctx context.Context, ad *ad.Entity) error
	GetByID(ctx context.Context, id string) (*ad.Entity, error)
	GetAll(ctx context.Context) ([]*ad.Entity, error)
	Update(ctx context.Context, ad *ad.Entity) error
	Delete(ctx context.Context, id string) error
	GetByInfluencerID(ctx context.Context, influencerID string) ([]*ad.Entity, error)
	GetByCompanyName(ctx context.Context, companyName string) ([]*ad.Entity, error)
}

type Company interface {
	Create(ctx context.Context, company *company.Entity) (string, error)
	GetByEmail(ctx context.Context, email string) (*company.Entity, error)
	GetByID(ctx context.Context, id string) (*company.Entity, error)
	Update(ctx context.Context, company *company.Entity) error
	Delete(ctx context.Context, id string) error
}

type Influencer interface {
	Create(ctx context.Context, influencer *influencer.Entity) (string, error)
	GetByEmail(ctx context.Context, email string) (*influencer.Entity, error)
	GetByID(ctx context.Context, id string) (*influencer.Entity, error)
	GetAll(ctx context.Context) ([]*influencer.Entity, error)
	Update(ctx context.Context, influencer *influencer.Entity) error
	Delete(ctx context.Context, id string) error
}

type AdResponse interface {
	Create(ctx context.Context, adResponse *ad_response.CreateRequest) error
}

type Repositories struct {
	Company
	Influencer
	Ad
	AdResponse
}

func NewRepositories(pgx *pgxpool.Pool) *Repositories {
	return &Repositories{
		Company:    postgresql.NewCompanyRepo(pgx),
		Influencer: postgresql.NewInfluencerRepo(pgx),
		Ad:         postgresql.NewAdRepo(pgx),
		AdResponse: postgresql.NewAdResponseRepo(pgx),
	}
}
