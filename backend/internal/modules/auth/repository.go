package auth

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var errUserNotFound = errors.New("user not found")

type userRow struct {
	ID           string
	Username     string
	PasswordHash string
	Role         string
	DisplayName  string
}

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) FindByUsername(ctx context.Context, username string) (*userRow, error) {
	var u userRow
	err := r.pool.QueryRow(ctx, `
		SELECT id, username, password_hash, role, display_name
		FROM users
		WHERE username = $1 AND active = TRUE
	`, username).Scan(&u.ID, &u.Username, &u.PasswordHash, &u.Role, &u.DisplayName)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errUserNotFound
		}
		return nil, err
	}
	return &u, nil
}
