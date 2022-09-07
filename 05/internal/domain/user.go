package domain

import "context"

type User struct {
	ID        int
	Username  string
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type UserRepo interface {
	New(ctx context.Context, user *User) error
	Get(ctx context.Context, userID int) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	SearchByEmail(ctx context.Context, email string) (*User, error)
	Delete(ctx context.Context, userID int) error
}
