package repo

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/wuzyk/otus-microservice-arch/02/internal/domain"
)

type UserRow struct {
	bun.BaseModel `bun:"table:users"`

	ID        int    `bun:"id,pk,autoincrement"`
	Username  string `bun:"username"`
	FirstName string `bun:"firstname"`
	LastName  string `bun:"lastname"`
	Email     string `bun:"email"`
	Phone     string `bun:"phone"`
}

type UserRepo struct {
	db *bun.DB
}

func NewUserRepo(db *bun.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) New(ctx context.Context, user *domain.User) error {
	userRow := UserRow{
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}

	_, err := r.db.
		NewInsert().
		Model(&userRow).
		Exec(ctx)

	user.ID = userRow.ID

	return err
}

func (r *UserRepo) Get(ctx context.Context, userID int) (*domain.User, error) {
	var userRow UserRow

	err := r.db.
		NewSelect().
		Model(&userRow).
		Where("id = ?", userID).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        userRow.ID,
		Username:  userRow.Username,
		FirstName: userRow.FirstName,
		LastName:  userRow.LastName,
		Email:     userRow.Email,
		Phone:     userRow.Phone,
	}, nil
}

func (r *UserRepo) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	userRow := UserRow{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}

	_, err := r.db.
		NewUpdate().
		Model(&userRow).
		ExcludeColumn("id", "username").
		WherePK().
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        userRow.ID,
		Username:  userRow.Username,
		FirstName: userRow.FirstName,
		LastName:  userRow.LastName,
		Email:     userRow.Email,
		Phone:     userRow.Phone,
	}, nil
}

func (r *UserRepo) Delete(ctx context.Context, userID int) error {
	_, err := r.db.
		NewDelete().
		Table("users").
		Where("id = ?", userID).
		Exec(ctx)

	return err
}
