package postgresql

import (
	"Ads-marketplace/internal/domain/ad_response"
	"Ads-marketplace/internal/domain/influencer"
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

func (r *InfluencerRepo) Create(ctx context.Context, influencer *influencer.Entity) (string, error) {
	var id string

	err := r.db.QueryRow(ctx, `
		INSERT INTO influencers (id, name, email, password, phone, account_type, platforms)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`, influencer.ID, influencer.Name, influencer.Email, influencer.Password, influencer.Phone, influencer.AccountType, influencer.Platforms).Scan(&id)

	return id, err
}

func (r *InfluencerRepo) GetByEmail(ctx context.Context, email string) (*influencer.Entity, error) {
	var influencer influencer.Entity

	err := r.db.QueryRow(ctx, `
		SELECT id, name, email, password, phone, platforms, category, account_type
		FROM influencers WHERE email = $1
	`, email).Scan(&influencer.ID, &influencer.Name, &influencer.Email, &influencer.Password, &influencer.Phone, &influencer.Platforms, &influencer.Category, &influencer.AccountType)
	if err != nil {
		return nil, err
	}

	return &influencer, nil
}

func (r *InfluencerRepo) GetByID(ctx context.Context, id string) (*influencer.Entity, error) {
	row := r.db.QueryRow(ctx, `
		SELECT id, name, email, password, phone, platforms, category, account_type, created_at, updated_at
		FROM influencers
		WHERE id = $1
	`, id)

	var inf influencer.Entity
	err := row.Scan(
		&inf.ID, &inf.Name, &inf.Email, &inf.Password, &inf.Phone, &inf.Platforms, &inf.Category, &inf.AccountType, &inf.CreatedAt, &inf.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &inf, nil
}

func (r *InfluencerRepo) GetAll(ctx context.Context) ([]*influencer.Entity, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, name, email, password, phone, platforms, category, account_type, created_at, updated_at
		FROM influencers
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var influencers []*influencer.Entity

	for rows.Next() {
		var inf influencer.Entity
		var platforms []string

		err := rows.Scan(
			&inf.ID, &inf.Name, &inf.Email, &inf.Password, &inf.Phone, &platforms, &inf.Category, &inf.AccountType, &inf.CreatedAt, &inf.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		inf.Platforms = platforms
		influencers = append(influencers, &inf)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return influencers, nil
}

func (r *InfluencerRepo) Update(ctx context.Context, influencer *influencer.Entity) error {
	_, err := r.db.Exec(ctx, `
		UPDATE influencers
		SET name = $1, email = $2, phone = $3, platforms = $4, category = $5, city = $6, updated_at = NOW()
		WHERE id = $7
	`, influencer.Name, influencer.Email, influencer.Phone, influencer.Platforms, influencer.Category, influencer.City, influencer.ID)
	return err
}

func (r *InfluencerRepo) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM influencers WHERE id = $1`, id)
	return err
}

func (r *InfluencerRepo) GetAdsResponsesByID(ctx context.Context, influencerID string) ([]*ad_response.AdResponseDTO, error) {
	query := `
		SELECT 
			a.id, a.title, a.company_name,
			ar.id, ar.status
		FROM ad_responses ar
		JOIN ads a ON ar.ad_id = a.id
		WHERE ar.influencer_id = $1
	`

	rows, err := r.db.Query(ctx, query, influencerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*ad_response.AdResponseDTO
	for rows.Next() {
		var item ad_response.AdResponseDTO
		if err := rows.Scan(
			&item.AdsID, &item.Title, &item.CompanyName, &item.ID, &item.Status,
		); err != nil {
			return nil, err
		}
		results = append(results, &item)
	}

	return results, nil
}
