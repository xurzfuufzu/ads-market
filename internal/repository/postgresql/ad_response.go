package postgresql

import (
	"Ads-marketplace/internal/domain/ad_response"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AdResponseRepo struct {
	db *pgxpool.Pool
}

func NewAdResponseRepo(db *pgxpool.Pool) *AdResponseRepo {
	return &AdResponseRepo{db: db}
}

func (r *AdResponseRepo) Create(ctx context.Context, adResponse *ad_response.CreateRequest) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO ad_responses (ad_id, influencer_id, message)
		VALUES ($1, $2, $3)
	`, adResponse.AdID, adResponse.InfluencerID, adResponse.Message)

	if err != nil {
		return fmt.Errorf("failed to create ad response: %v", err)
	}

	_, err = r.db.Exec(ctx, `
		UPDATE ads
		SET responses_count = responses_count + 1
		WHERE id = $1
	`, adResponse.AdID)

	if err != nil {
		return fmt.Errorf("failed to update ad responses count: %v", err)
	}

	return nil
}

func (r *AdResponseRepo) UpdateStatus(ctx context.Context, dto ad_response.UpdateAdStatusDTO) error {
	res, err := r.db.Exec(ctx, `
		UPDATE ad_responses
		SET status = $1
		WHERE id = $2
	`, dto.Status, dto.ID)

	if err != nil {
		return fmt.Errorf("failed to update ad response status: %v", err)
	}

	rows := res.RowsAffected()
	if rows == 0 {
		fmt.Printf("No row found with id: %s\n", dto.ID)
	}

	return nil
}
