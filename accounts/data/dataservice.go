package data

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"time"

	"github.com/chutified/appointments/accounts/config"
	"github.com/chutified/appointments/accounts/models"
	"github.com/google/uuid"
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

// ErrScanRow is returned when the query returns unexpected result.
var ErrScanRow = errors.New("unexpected scan's destination")

// AddAccount adds a new account into the database and created the generated ID.
func (ds *DatabaseService) AddAccount(ctx context.Context, a *models.Account) (string, error) {

	// get the sql
	sqls, err := getQuery("add_account.sql")
	if err != nil {
		return "", errors.Wrap(err, "getting the sql")
	}

	// generate uuid
	id := uuid.New().String()

	// parse birth day
	y, m, d := a.BirthDay.Date()
	birthDay := fmt.Sprintf("%d-%d-%d", y, m, d)

	// run the sql
	_, err = ds.db.ExecContext(ctx, sqls, id, a.Username, a.Email, a.Phone, a.HPassword, a.FirstName, a.LastName, birthDay, a.PermanentAddress, a.MailingAddress)
	if err != nil {
		return "", errors.Wrap(err, "inserting a new user")
	}

	return id, nil
}

// AccountsPages return the number of pages with pageCap items on each page.
func (ds *DatabaseService) AccountsPages(ctx context.Context, pageCap int) (int, error) {

	// get sql
	sqls, err := getQuery("pages.sql")
	if err != nil {
		return 0, errors.Wrap(err, "getting the sql")
	}

	// run sql
	var l int
	err = ds.db.QueryRowContext(ctx, sqls).Scan(&l)
	if err != nil {
		return 0, errors.Wrap(err, "query for the number of rows")
	}

	// calculate the number of pages
	p := math.Ceil(float64(l) / float64(pageCap))
	pages := int(p)

	return pages, nil
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

// ErrSQLFileNotFound is returned when no sql file is found.
var ErrSQLFileNotFound = errors.New("the sql file was not found")

// getQuery reads the sql from the sql file and returns it in a string form.
func getQuery(file string) (string, error) {

	// read file
	path := filepath.Join("queries", file)
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return "", ErrSQLFileNotFound
	}

	return string(bs), nil
}
