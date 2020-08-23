package models

// Account defines all stored variables for an account
type Account struct {
	ID        string `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	HPassword string `json:"hashed_password,omitempty"`

	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`

	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	BirthDay  string `json:"birth_day,omitempty"`

	PermanentAddress string `json:"permanent_address,omitempty"`
	MailingAddress   string `json:"mailing_address,omitempty"`

	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty"`
}

// Login is the model of the Login/Validate request.
type Login struct {
	ID        string `json:"id,omitempty"`
	HPassword string `json:"hashed_password,omitempty"`
}
