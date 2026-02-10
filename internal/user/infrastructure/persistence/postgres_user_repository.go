package persistence

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/naeemjr06-prog/user-access-api/internal/user/domain"
)

type PostgresUserRepository struct {
	db *pgxpool.Pool
}

func NewPostgresUserRepository(db *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(ctx context.Context, u *domain.User) error {
	u.ID = uuid.NewString()

	_, err := r.db.Exec(ctx, `
		INSERT INTO users (id, email, password_hash, role, active, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
	`,
		u.ID, u.Email, u.Password, string(u.Role), u.Active, u.CreatedAt, u.UpdatedAt,
	)

	return err
}

func (r *PostgresUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	row := r.db.QueryRow(ctx, `
		SELECT id, email, password_hash, role, active, created_at, updated_at
		FROM users WHERE email=$1 LIMIT 1
	`, email)

	var u domain.User
	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.Role, &u.Active, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, errors.New("not found")
	}

	return &u, nil
}
