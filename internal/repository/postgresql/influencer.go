package postgresql

import (
	"Ads-marketplace/internal/domain/influencer"
	"Ads-marketplace/pkg/utils"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type InfluencerRepo struct {
	db *pgxpool.Pool
}

func NewInfluencerRepo(db *pgxpool.Pool) *InfluencerRepo {
	return &InfluencerRepo{
		db: db,
	}
}

func (r *InfluencerRepo) Create(ctx context.Context, influencer *influencer.Entity) error {
	platformsJson, err := utils.SerializePlatforms(influencer.Platforms)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, `
		INSERT INTO influencers (id, name, email, password, phone, platforms, account_type)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, influencer.ID, influencer.Name, influencer.Email, influencer.Password, influencer.Phone, platformsJson, influencer.AccountType)
	return err
}

func (r *InfluencerRepo) GetByEmail(ctx context.Context, email string) (*influencer.Entity, error) {
	var influencer influencer.Entity
	var platformsJson string

	err := r.db.QueryRow(ctx, `
		SELECT id, name, email, password, phone, platforms, account_type
		FROM influencers WHERE email = $1
	`, email).Scan(&influencer.ID, &influencer.Name, &influencer.Email, &influencer.Password, &influencer.Phone, &platformsJson, &influencer.AccountType)
	if err != nil {
		return nil, err
	}

	platforms, err := utils.DeserializePlatforms(platformsJson)
	if err != nil {
		return nil, err
	}

	influencer.Platforms = platforms
	return &influencer, nil
}

func (r *InfluencerRepo) Update(ctx context.Context, influencer *influencer.Entity) error {
	platformsJson, err := utils.SerializePlatforms(influencer.Platforms)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, `
		UPDATE influencers SET name = $1, email = $2, password = $3, phone = $4, platforms = $5, account_type = $6, updated_at = $7
		WHERE id = $8
	`, influencer.Name, influencer.Email, influencer.Password, influencer.Phone, platformsJson, influencer.AccountType, influencer.UpdatedAt, influencer.ID)
	return err
}

func (r *InfluencerRepo) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM influencers WHERE id = $1`, id)
	return err
}
