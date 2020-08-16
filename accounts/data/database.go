package data

import (
	"context"

	"github.com/chutified/appointments/accounts/models"
)

// DatabaseService is an interface, which is able to make the database actions.
type DatabaseService interface {

	// Create
	AddAccount(ctx context.Context, a *models.Account) (*models.Account, error)

	// Read
	GetAccountsAll(ctx context.Context, pageCap int, pageNum int) ([]*models.Account, error)
	GetAccountByID(ctx context.Context, id string) (*models.Account, error)
	GetAccountByParams(ctx context.Context, a *models.Account) (*models.Account, error)
	LoginAccount(ctx context.Context, email string, hPasswd string) (*models.Account, error)

	// Update
	EditAccountByID(ctx context.Context, id string) (*models.Account, error)

	// Delete
	DeleteAccountByID(ctx context.Context, id string) (*models.Account, error)
}
