package models

// Account defines all stored variables for an account
type Account struct {
	ID        string
	Username  string
	HPassword string

	Email string
	Phone string

	FirstName string
	LastName  string
	BirthDay  string

	PermanentAddress string
	MailingAddress   string

	CreatedAt string
	UpdatedAt string
	DeletedAt string
}
