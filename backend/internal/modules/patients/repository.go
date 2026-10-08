package patients

import (
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var errPatientNotFound = errors.New("patient not found")

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) List(ctx context.Context, q string) ([]Patient, error) {
	q = strings.TrimSpace(q)
	rows, err := r.pool.Query(ctx, `
		SELECT id, code, full_name, mobile, COALESCE(age, 0), COALESCE(gender, ''),
		       visits, invoices, total_paid, wallet, status
		FROM patients
		WHERE ($1 = '' OR full_name ILIKE '%' || $1 || '%' OR mobile ILIKE '%' || $1 || '%' OR code ILIKE '%' || $1 || '%')
		ORDER BY full_name
	`, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Patient
	for rows.Next() {
		var p Patient
		if err := rows.Scan(&p.ID, &p.Code, &p.FullName, &p.Mobile, &p.Age, &p.Gender,
			&p.Visits, &p.Invoices, &p.TotalPaid, &p.Wallet, &p.Status); err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, rows.Err()
}

func (r *Repository) GetByID(ctx context.Context, id string) (*PatientProfile, error) {
	var p PatientProfile
	err := r.pool.QueryRow(ctx, `
		SELECT id, code, full_name, mobile, COALESCE(age, 0), COALESCE(gender, ''),
		       visits, invoices, total_paid, wallet, status, COALESCE(notes, '')
		FROM patients WHERE id = $1
	`, id).Scan(
		&p.ID, &p.Code, &p.FullName, &p.Mobile, &p.Age, &p.Gender,
		&p.Visits, &p.Invoices, &p.TotalPaid, &p.Wallet, &p.Status, &p.Notes,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errPatientNotFound
		}
		return nil, err
	}
	return &p, nil
}

func (r *Repository) WalletBalance(ctx context.Context, id string) (int64, error) {
	var balance int64
	err := r.pool.QueryRow(ctx, `SELECT wallet FROM patients WHERE id = $1`, id).Scan(&balance)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, errPatientNotFound
		}
		return 0, err
	}
	return balance, nil
}
