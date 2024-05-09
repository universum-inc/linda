package account

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetAccounts(limit int64) ([]Account, error)
	GetAccountByID(id int64) (*Account, error)
	GetAccountByUsername(username string) (*Account, error)
}

type RepositoryImpl struct {
	DB *sqlx.DB
}

func NewRepository(dataSourceName string) (*RepositoryImpl, error) {
	db, err := sqlx.Connect("pgx", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &RepositoryImpl{DB: db}, nil
}

func (repo *RepositoryImpl) GetAccounts(limit int64) ([]Account, error) {
	var accounts []Account
	err := repo.DB.Select(&accounts, "SELECT * FROM account LIMIT $1", limit)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (repo *RepositoryImpl) GetAccountByID(id int64) (*Account, error) {
	var account Account
	err := repo.DB.Get(&account, "SELECT * FROM account WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (repo *RepositoryImpl) GetAccountByUsername(username string) (*Account, error) {
	var account Account
	err := repo.DB.Get(&account, "SELECT * FROM account WHERE username = $1", username)
	if err != nil {
		return nil, err
	}
	return &account, nil
}
