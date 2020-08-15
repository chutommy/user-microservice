package data

import (
	"context"
	"database/sql"

	"github.com/chutified/appointments/accounts/models"
)

// Database is an interface, which is able to make the database actions.
type Database interface {

	// Create
	AddAccount(ctx context.Context, u *models.Account)

	// Read
	GetAccountsAll(ctx context.Context, pageCap int, pageNum int)
	GetAccountByID(ctx context.Context, id string)
	GetAccountByParams(ctx context.Context, u *models.Account)
	LoginAccount(ctx context.Context, email string, hPasswd string)

	// Update
	EditAccountById(ctx context.Context, id string)

	// Delete
	DeleteAccountById(ctx context.Context, id string)
}

type databaseService struct {
	db *sql.DB
}

func New() *databaseService {
	return &databaseService{}
}
