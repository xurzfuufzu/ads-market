package postgresql

import (
	"Ads-marketplace/internal/domain/ad"
	"Ads-marketplace/internal/domain/company"
	"Ads-marketplace/internal/domain/influencer"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CompanyRepo struct {
	db *pgxpool.Pool
}

func NewCompanyRepo(db *pgxpool.Pool) *CompanyRepo {
	return &CompanyRepo{
		db: db,
	}
}

func (r *CompanyRepo) Create(ctx context.Context, company *company.Entity) (string, error) {
	var id string

	err := r.db.QueryRow(ctx, `
		INSERT INTO companies (id, name, email, password, phone, account_type)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`, company.ID, company.Name, company.Email, company.Password, company.Phone, company.AccountType).Scan(&id)

	return id, err
}

func (r *CompanyRepo) GetByEmail(ctx context.Context, email string) (*company.Entity, error) {
	var company company.Entity
	err := r.db.QueryRow(ctx, `
    SELECT id, name, email, password, phone, description, account_type
    FROM companies WHERE email ILIKE $1
`, email).Scan(&company.ID, &company.Name, &company.Email, &company.Password, &company.Phone, &company.Description, &company.AccountType)

	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *CompanyRepo) GetByID(ctx context.Context, id string) (*company.Entity, error) {
	row := r.db.QueryRow(ctx, `
		SELECT id, name, email, password, phone, description, account_type, created_at, updated_at
		FROM companies
		WHERE id = $1
	`, id)

	var c company.Entity
	err := row.Scan(
		&c.ID, &c.Name, &c.Email, &c.Password, &c.Phone, &c.Description, &c.AccountType, &c.CreatedAt, &c.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *CompanyRepo) GetCompanyAds(ctx context.Context, companyName string) ([]*ad.Entity, error) {
	query := `
		SELECT id, title, company_name, description, pricefrom, priceto, category, target_city, platforms, status, created_at, updated_at, responses_count
		FROM ads
		WHERE company_name = $1
	`

	rows, err := r.db.Query(ctx, query, companyName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ads []*ad.Entity
	for rows.Next() {
		var a ad.Entity
		if err := rows.Scan(
			&a.ID, &a.Title, &a.CompanyName, &a.Description, &a.PriceFrom, &a.PriceTo, &a.Category,
			&a.City, &a.Platforms, &a.Status, &a.CreatedAt, &a.UpdatedAt, &a.ResponsesCount,
		); err != nil {
			return nil, err
		}
		ads = append(ads, &a)
	}

	return ads, nil
}

func (r *CompanyRepo) Update(ctx context.Context, company *company.Entity) error {
	_, err := r.db.Exec(ctx, `
		UPDATE companies
		SET name = $1, email = $2, phone = $3, description = $4, updated_at = NOW()
		WHERE id = $5
	`, company.Name, company.Email, company.Phone, company.Description, company.ID)
	return err
}

func (r *CompanyRepo) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM companies WHERE id = $1`, id)
	return err
}

func (r *CompanyRepo) GetInfluencersForAd(ctx context.Context, id uuid.UUID) ([]*influencer.InfluencerDTO, error) {
	query := `
		SELECT
    		i.id, i.name,i.platforms,i.category,i.city,ar.status
			FROM ad_responses ar
         	JOIN ads a ON ar.ad_id = a.id
         	JOIN influencers i ON ar.influencer_id = i.id
			WHERE a.id = $1
	`

	fmt.Println(id)

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch influencers for ad: %v", err)
	}
	defer rows.Close()

	var results []*influencer.InfluencerDTO
	for rows.Next() {
		var item influencer.InfluencerDTO
		if err := rows.Scan(
			&item.ID, &item.Name, &item.Category, &item.Platforms,
			&item.City, &item.Status,
		); err != nil {
			return nil, fmt.Errorf("failed to scan influencer data: %v", err)
		}
		results = append(results, &item)
	}

	return results, nil
}
