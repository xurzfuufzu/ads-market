package postgresql

import (
	"Ads-marketplace/internal/domain/ad"
	"Ads-marketplace/pkg/utils"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AdRepo struct {
	db *pgxpool.Pool
}

func NewAdRepo(db *pgxpool.Pool) *AdRepo {
	return &AdRepo{db: db}
}

func (r *AdRepo) Create(ctx context.Context, a *ad.Entity) error {
	platformsJSON, err := utils.SerializePlatforms(a.Platform)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, `
		INSERT INTO ads (id, title, company_name, influencer_id, description, price, status, created_at, updated_at, platforms, category, target_country)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`,
		a.ID, a.Title, a.CompanyName, a.InfluencerId, a.Description,
		a.Price, a.Status, a.CreatedAt, a.UpdatedAt, platformsJSON, a.Category, a.TargetCountry,
	)

	return err
}

func (r *AdRepo) GetByID(ctx context.Context, id string) (*ad.Entity, error) {
	var a ad.Entity
	var platformsJSON string
	var influencerID *string

	err := r.db.QueryRow(ctx, `
		SELECT id, title, company_name, influencer_id, description, price, status, created_at, updated_at, platforms, category, target_country
		FROM ads WHERE id = $1
	`, id).Scan(
		&a.ID, &a.Title, &a.CompanyName, &influencerID, &a.Description, &a.Price, &a.Status,
		&a.CreatedAt, &a.UpdatedAt, &platformsJSON, &a.Category, &a.TargetCountry,
	)

	if err != nil {
		return nil, err
	}

	a.Platform, err = utils.DeserializePlatforms(platformsJSON)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (r *AdRepo) GetAll(ctx context.Context) ([]*ad.Entity, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, title, company_name, influencer_id, description, price, status, created_at, updated_at, platforms, category, target_country
		FROM ads
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ads []*ad.Entity
	for rows.Next() {
		var a ad.Entity
		var platformsJSON string
		var influencerID *string

		err := rows.Scan(
			&a.ID, &a.Title, &a.CompanyName, &influencerID, &a.Description, &a.Price, &a.Status,
			&a.CreatedAt, &a.UpdatedAt, &platformsJSON, &a.Category, &a.TargetCountry,
		)
		if err != nil {
			return nil, err
		}

		platforms, err := utils.DeserializePlatforms(platformsJSON)
		if err != nil {
			return nil, err
		}
		a.Platform = platforms

		ads = append(ads, &a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ads, nil
}

func (r *AdRepo) Update(ctx context.Context, a *ad.Entity) error {
	platformsJSON, err := utils.SerializePlatforms(a.Platform)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, `
		UPDATE ads SET title = $1, company_name = $2, influencer_id = $3, description = $4, price = $5,
			status = $6, updated_at = $7, platforms = $8, category = $9, target_country = $10
		WHERE id = $11
	`, a.Title, a.CompanyName, a.InfluencerId, a.Description, a.Price, a.Status,
		a.UpdatedAt, platformsJSON, a.Category, a.TargetCountry, a.ID)

	return err
}

func (r *AdRepo) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM ads WHERE id = $1`, id)
	return err
}

func (r *AdRepo) GetByInfluencerID(ctx context.Context, influencerID string) ([]*ad.Entity, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, title, company_name, influencer_id, description, price, status, created_at, updated_at, category, target_country
		FROM ads WHERE influencer_id = $1
	`, influencerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ads []*ad.Entity
	for rows.Next() {
		var a ad.Entity
		var platformsJSON string
		var influencerID *string

		err := rows.Scan(
			&a.ID, &a.Title, &a.CompanyName, &influencerID, &a.Description, &a.Price, &a.Status,
			&a.CreatedAt, &a.UpdatedAt, &platformsJSON, &a.Category, &a.TargetCountry,
		)
		if err != nil {
			return nil, err
		}

		platforms, err := utils.DeserializePlatforms(platformsJSON)
		if err != nil {
			return nil, err
		}
		a.Platform = platforms

		ads = append(ads, &a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ads, nil
}

func (r *AdRepo) GetByCompanyName(ctx context.Context, companyName string) ([]*ad.Entity, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, title, company_name, influencer_id, description, price, status, created_at, updated_at, category, target_country
		FROM ads WHERE company_name = $1
	`, companyName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ads []*ad.Entity
	for rows.Next() {
		var a ad.Entity
		var platformsJSON string
		var influencerID *string

		err := rows.Scan(
			&a.ID, &a.Title, &a.CompanyName, &influencerID, &a.Description, &a.Price, &a.Status,
			&a.CreatedAt, &a.UpdatedAt, &platformsJSON, &a.Category, &a.TargetCountry,
		)
		if err != nil {
			return nil, err
		}

		platforms, err := utils.DeserializePlatforms(platformsJSON)
		if err != nil {
			return nil, err
		}
		a.Platform = platforms

		ads = append(ads, &a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ads, nil
}
