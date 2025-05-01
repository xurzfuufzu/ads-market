package postgresql

import (
	"Ads-marketplace/internal/domain/company"
	"context"
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

func (r *CompanyRepo) Update(ctx context.Context, company *company.Entity) error {
	_, err := r.db.Exec(ctx, `
		UPDATE companies SET name = $1, email = $2, password = $3, phone = $4, description = $5, account_type = $6, updated_at = $7
		WHERE id = $8
	`, company.Name, company.Email, company.Password, company.Phone, company.Description, company.AccountType, company.UpdatedAt, company.ID)
	return err
}

func (r *CompanyRepo) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM companies WHERE id = $1`, id)
	return err
}
