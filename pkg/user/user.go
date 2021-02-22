package user

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/upper/db/v4/adapter/cockroachdb"
)

type User struct {
	ID              string          `db:"id,omitempty"`
	CreatedAt       time.Time       `db:"created_at"`
	UpdatedAt       time.Time       `db:"updated_at"`
	LastName        string          `db:"last_name"`
	FirstName       string          `db:"first_name"`
	NickName        string          `db:"nick_name"`
	DateOfBirth     string          `db:"date_of_birth"`
	DriversLicense  *DriversLicense `db:"drivers_license"`
	Passport        *Passport       `db:"passport"`
	DriversLicense2 *DriversLicense `db:"drivers_license2"`
	Passport2       *Passport       `db:"passport2"`
	CreditCards     []*CreditCard   `gorm:"-"`
}

type DriversLicense struct {
	ID       string
	Class    string
	ExpireAt time.Time
	Donor    bool
	*cockroachdb.JSONBConverter
}

type Passport struct {
	Number    string
	Country   string
	Signature []byte
	*cockroachdb.JSONBConverter
}

type CreditCard struct {
	CardNumber   string
	SecurityCode int
}

func (d *DriversLicense) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	err := json.Unmarshal(bytes, d)
	return err
}

func (d *DriversLicense) Value() (driver.Value, error) {
	return json.Marshal(d)
}

func (p *Passport) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	err := json.Unmarshal(bytes, p)
	return err
}

func (p *Passport) Value() (driver.Value, error) {
	return json.Marshal(p)
}
