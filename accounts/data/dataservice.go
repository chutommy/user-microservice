package data

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/chutified/appointments/accounts/config"
	"github.com/chutified/appointments/accounts/models"
	"github.com/pkg/errors"
)

// DatabaseService manages all database operations.
type DatabaseService struct {
	db *sql.DB
}

// New is the contructor for the DatabaseService.
func New() *DatabaseService {
	return &DatabaseService{}
}

// Init initialize the DatabaseService connection to the database.
func (ds *DatabaseService) Init(cfg *config.DBConfig) error {

	// define database connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBname)

	// open database connection
	var err error
	ds.db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return errors.Wrap(err, "connecting to the db")
	}

	// test the connection
	for i := 0; i < 3; i++ {
		err = ds.db.Ping()
		if err == nil {
			break
		}

		time.Sleep(3 * time.Second)
	}
	if err != nil {
		return errors.Wrap(err, "db connection failed")
	}

	return nil
}

// Stop close the database connection.
func (ds *DatabaseService) Stop() error {

	// close database connection
	err := ds.db.Close()
	if err != nil {
		return errors.Wrap(err, "stoping database service")
	}

	return nil
}

func (ds *DatabaseService) AddAccount(ctx context.Context, a *models.Account) (*models.Account, error) {

	return nil, nil
}
func (ds *DatabaseService) GetAccountsAll(ctx context.Context, pageCap int, pageNum int) ([]*models.Account, error) {

	return nil, nil
}
func (ds *DatabaseService) GetAccountByID(ctx context.Context, id string) (*models.Account, error) {

	return nil, nil
}
func (ds *DatabaseService) GetAccountByParams(ctx context.Context, a *models.Account) (*models.Account, error) {

	return nil, nil
}
func (ds *DatabaseService) LoginAccount(ctx context.Context, email string, hPasswd string) (*models.Account, error) {

	return nil, nil
}
func (ds *DatabaseService) EditAccountByID(ctx context.Context, id string) (*models.Account, error) {

	return nil, nil
}
func (ds *DatabaseService) DeleteAccountByID(ctx context.Context, id string) (*models.Account, error) {

	return nil, nil
}
