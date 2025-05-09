package postgresql

import (
	"Ads-marketplace/internal/domain/ad"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AdRepo struct {
	db *pgxpool.Pool
}

func NewAdRepo(db *pgxpool.Pool) *AdRepo {
	return &AdRepo{db: db}
}

func (r *AdRepo) Create(ctx context.Context, a *ad.Entity) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO ads (id, title, company_name, description, priceFrom, priceTo, status, created_at, updated_at, platforms, category, target_city)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`, a.ID, a.Title, a.CompanyName, a.Description, a.PriceFrom, a.PriceTo, a.Status, a.CreatedAt, a.UpdatedAt, a.Platforms, a.Category, a.City)

	return err
}

func (r *AdRepo) GetAll(ctx context.Context) ([]*ad.Entity, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, title, company_name, description, priceFrom, priceTo, status, created_at, updated_at, platforms, category, target_city, responses_count
		FROM ads
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch ads: %v", err)
	}
	defer rows.Close()

	var ads []*ad.Entity
	for rows.Next() {
		var a ad.Entity

		err := rows.Scan(
			&a.ID, &a.Title, &a.CompanyName, &a.Description, &a.PriceFrom, &a.PriceTo, &a.Status, &a.CreatedAt, &a.UpdatedAt, &a.Platforms, &a.Category, &a.City, &a.ResponsesCount,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan ad: %v", err)
		}

		ads = append(ads, &a)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return ads, nil
}

func (r *AdRepo) GetByID(ctx context.Context, id string) (*ad.Entity, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdRepo) Update(ctx context.Context, ad *ad.Entity) error {
	_, err := r.db.Exec(ctx, `
		UPDATE ads
		SET title = $1, company_name = $2, description = $3, pricefrom = $4, priceto = $5, 
		    status = $6, platforms = $7, category = $8,target_city = $9,updated_at = NOW()
		WHERE id = $10
	`, ad.Title, ad.CompanyName, ad.Description, ad.PriceFrom, ad.PriceTo, ad.Status, ad.Platforms, ad.Category, ad.City, ad.ID)

	return err
}

func (r *AdRepo) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM ads WHERE id = $1`, id)
	return err
}

func (r *AdRepo) GetByInfluencerID(ctx context.Context, influencerID string) ([]*ad.Entity, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AdRepo) GetByCompanyName(ctx context.Context, companyName string) ([]*ad.Entity, error) {
	//TODO implement me
	panic("implement me")
}
