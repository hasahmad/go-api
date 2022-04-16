package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hasahmad/go-api/pkg/validator"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4"
)

var AnonymousUser = &User{}

type User struct {
	UserID      uuid.UUID   `db:"user_id" json:"user_id" goqu:"defaultifempty,skipupdate"`
	FirstName   null.String `db:"first_name" json:"first_name" goqu:"defaultifempty"`
	LastName    null.String `db:"last_name" json:"last_name" goqu:"defaultifempty"`
	Username    string      `db:"username" json:"username" goqu:"skipupdate"`
	Email       null.String `db:"email" json:"email" goqu:"defaultifempty"`
	Password    password    `db:"password" json:"-"`
	IsStaff     bool        `db:"is_staff" json:"is_staff" goqu:"defaultifempty"`
	IsSuperuser bool        `db:"is_superuser" json:"is_superuser" goqu:"defaultifempty"`
	LastLogin   null.Time   `db:"last_login" json:"last_login"`
	Version     int         `db:"version" json:"version"`
	CreatedAt   time.Time   `db:"created_at" json:"created_at" goqu:"defaultifempty,skipupdate"`
	UpdatedAt   time.Time   `db:"updated_at" json:"updated_at" goqu:"defaultifempty"`
	DeletedAt   null.Time   `db:"deleted_at" json:"deleted_at"`
	Roles       []Role      `db:"-" json:"roles,omitempty"`
}

func (u *User) IsAnonymousUser() bool {
	return u == AnonymousUser
}

func NewUser(username string, pass string) (*User, error) {
	p := password{}
	err := p.Set(pass)
	if err != nil {
		return nil, err
	}

	return &User{
		UserID:      uuid.New(),
		FirstName:   null.StringFromPtr(nil),
		LastName:    null.StringFromPtr(nil),
		Username:    username,
		Email:       null.StringFromPtr(nil),
		Password:    p,
		IsStaff:     false,
		IsSuperuser: false,
		Version:     1,
		LastLogin:   null.TimeFromPtr(nil),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

type password struct {
	plaintext *string
	hash      []byte
}

func (p *password) Scan(value interface{}) error {
	if value == nil {
		p.plaintext, p.hash = nil, nil
		return nil
	}
	p.plaintext = nil
	v, ok := value.([]byte)
	if !ok {
		// most likely a string
		vstr, ok := value.(string)
		if !ok {
			return fmt.Errorf("unable to convert password hash")
		} else {
			p.hash = []byte(vstr)
		}
	} else {
		p.hash = v
	}

	return nil
}

func (p password) Value() (driver.Value, error) {
	if p.hash == nil {
		return nil, nil
	}
	return p.hash, nil
}

func (p *password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}

	p.plaintext = &plaintextPassword
	p.hash = hash

	return nil
}

func (p *password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long")
}
