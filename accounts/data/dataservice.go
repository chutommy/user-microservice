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

// ErrQuery is returned when the query failed to execute p9roperly.
var ErrQuery = errors.New("query error")

// ErrExecuteSQL is returned when the SQL query failed to execute.
var ErrExecuteSQL = errors.New("failed to execute query")

// AddAccount adds a new account into the database and created the generated ID.
func (ds *DatabaseService) AddAccount(ctx context.Context, a *models.Account) (string, error) {

	// get the sql
	sqls, err := getQuery("add_account.sql")
	if err != nil {
		return "", errors.Wrap(err, "getting the sql")
	}

	// generate uuid
	id := uuid.New().String()

	// run the sql
	_, err = ds.db.ExecContext(ctx, sqls, id, a.Username, a.Email, a.Phone, a.HPassword, a.FirstName, a.LastName, a.BirthDay, a.PermanentAddress, a.MailingAddress)
	if err != nil {
		return "", errors.Wrap(ErrExecuteSQL, err.Error())
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
	if err == sql.ErrNoRows {
		return 0, err
	} else if err != nil {
		return 0, errors.Wrap(ErrQuery, err.Error())
	}

	// calculate the number of pages
	p := math.Ceil(float64(l) / float64(pageCap))
	pages := int(p)

	return pages, nil
}

// GetAllAccounts return all accounts in the detabase, except for deleted one.
func (ds *DatabaseService) GetAllAccounts(ctx context.Context, pageCap int, pageNum int, sortBy string, asc bool) ([]*models.Account, error) {

	// get sql
	sqls, err := getQuery("get_accounts.sql")
	if err != nil {
		return nil, errors.Wrap(err, "getting the sql")
	}

	// get the pagination
	offset := pageCap * (pageNum - 1)

	// the the direction
	var direct string
	if asc {
		direct = "ASC"
	} else {
		direct = "DESC"
	}

	// run sql
	rows, err := ds.db.QueryContext(ctx, sqls, pageCap, offset, sortBy, direct)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, errors.Wrap(ErrQuery, err.Error())
	}

	// prepare the array
	accounts := []*models.Account{}
	// range over rows
	for rows.Next() {

		// init th =e model
		var a *models.Account

		// scan the values into the model
		err := rows.Scan(&a.ID, &a.Username, &a.Email, &a.Phone, &a.FirstName, &a.LastName, &a.BirthDay, &a.PermanentAddress, &a.MailingAddress, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return nil, errors.Wrap(ErrScanRow, err.Error())
		}

		accounts = append(accounts, a)
	}

	return accounts, nil
}

// GetAccountByID returns an account with the given ID, returns nil if the account with the ID is not find or deleted.
func (ds *DatabaseService) GetAccountByID(ctx context.Context, id string) (*models.Account, error) {

	// get sql
	sqls, err := getQuery("get_by_id.sql")
	if err != nil {
		return nil, errors.Wrap(err, "getting the sql")
	}

	var a *models.Account
	// run sql and scan
	err = ds.db.QueryRowContext(ctx, sqls, id).Scan(&a.ID, &a.Username, &a.Email, &a.Phone, &a.FirstName, &a.LastName, &a.BirthDay, &a.PermanentAddress, &a.MailingAddress, &a.CreatedAt, &a.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, errors.Wrap(ErrQuery, err.Error())
	}

	return a, nil
}

// GetAccountByParams returns the account queried by non-nil parameters.
// TODO: handles multiple accounts for return
func (ds *DatabaseService) GetAccountByParams(ctx context.Context, a *models.Account) (*models.Account, error) {

	// get sql
	sqls, err := getQuery("get_by_params.sql")
	if err != nil {
		return nil, errors.Wrap(err, "getting the sql")
	}

	// run sql
	row := ds.db.QueryRowContext(ctx, sqls, a.ID, a.Username, a.Email, a.Phone, a.FirstName, a.LastName, a.BirthDay, a.PermanentAddress, a.MailingAddress, a.CreatedAt, a.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, errors.Wrap(ErrQuery, err.Error())
	}

	// parse the row
	var result *models.Account
	err = row.Scan(&result.ID, &result.Username, &result.Email, &result.Phone, &result.FirstName, &result.LastName, &result.BirthDay, &result.PermanentAddress, &result.MailingAddress, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, errors.Wrap(ErrScanRow, err.Error())
	}

	return result, nil
}

// ValidatePassword compares the given hashed password with the password of the the account.
// The True is returned if the passwords are same.
func (ds *DatabaseService) ValidatePassword(ctx context.Context, id string, hPasswd string) (bool, error) {

	// get sql
	sqls, err := getQuery("validate.sql")
	if err != nil {
		return false, errors.Wrap(err, "getting the sql")
	}

	var dbPasswd string
	// run sql
	err = ds.db.QueryRowContext(ctx, sqls, id).Scan(&dbPasswd)
	if err == sql.ErrNoRows {
		return false, err
	} else if err != nil {
		return false, errors.Wrap(ErrQuery, err.Error())
	}

	// compare
	ok := dbPasswd == hPasswd

	return ok, nil
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
