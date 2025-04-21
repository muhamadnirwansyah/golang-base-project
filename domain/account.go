package domain

import (
	"context"
)

type Account struct {
	ID          int64  `db:"id"`
	FullName    string `db:"full_name"`
	Email       string `db:"email"`
	PhoneNumber string `db:"phone_number"`
	Password    string `db:"password"`
}

type AccountRepository interface {
	Save(ctx context.Context, account *Account) error
	Update(ctx context.Context, account *Account) error
	FindById(ctx context.Context, id int64) (Account, error)
	FindByEmail(ctx context.Context, email string) (Account, error)
}
