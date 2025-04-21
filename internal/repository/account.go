package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/muhamadnirwansyah/authentication-service/domain"
)

type accountRepository struct {
	db *goqu.Database
}

func NewAccount(db *sql.DB) domain.AccountRepository {
	return &accountRepository{
		db: goqu.New("default", db),
	}
}

func (a *accountRepository) FindByEmail(ctx context.Context, email string) (account domain.Account, err error) {
	dataset := a.db.From("accounts").Where(goqu.Ex{
		"email": email,
	})
	_, err = dataset.ScanStructContext(ctx, &account)
	return
}

func (a *accountRepository) FindById(ctx context.Context, id int64) (account domain.Account, err error) {
	dataset := a.db.From("accounts").Where(goqu.Ex{
		"id": id,
	})
	_, err = dataset.ScanStructContext(ctx, &account)
	return
}

func (a *accountRepository) Save(ctx context.Context, account *domain.Account) error {
	var nextSequence int64

	err := a.db.QueryRowContext(ctx, "SELECT nextval('accounts_id_seq')").Scan(&nextSequence)
	if err != nil {
		return err
	}

	account.ID = nextSequence
	execute := a.db.Insert("accounts").Rows(account).Executor()
	_, err = execute.ExecContext(ctx)
	return err
}

func (a *accountRepository) Update(ctx context.Context, account *domain.Account) error {
	execute := a.db.Update("accounts").Set(goqu.Record{
		"email":        account.Email,
		"full_name":    account.FullName,
		"phone_number": account.PhoneNumber,
		"password":     account.Password,
	}).Where(goqu.Ex{"id": account.ID}).Executor()
	_, err := execute.ExecContext(ctx)
	return err
}
