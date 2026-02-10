package domain

import "context"

// Repository defines what persistence operations the domain needs.
// This is a PORT in Clean Architecture.
// Infrastructure (Postgres, Mongo, etc.) must implement this.
type Repository interface {
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
}
