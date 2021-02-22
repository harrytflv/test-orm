package userinline

import (
	"time"

	"github.com/harrytflv/test-orm/pkg/user"
)

type User struct {
	ID          string    `db:"id,omitempty"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	LastName    string    `db:"last_name"`
	FirstName   string    `db:"first_name"`
	NickName    string    `db:"nick_name"`
	DateOfBirth string    `db:"date_of_birth"`
	DID         string    `db:"d_id"`
	DClass      string    `db:"d_class"`
	DExpireAt   time.Time `db:"d_expire_at"`
	DDonor      bool      `db:"d_donor"`
	PNumber     string    `db:"p_number"`
	PCountry    string    `db:"p_country"`
	PSignature  []byte    `db:"p_signature"`
	DID2        string    `db:"d2_id"`
	DClass2     string    `db:"d2_class"`
	DExpireAt2  time.Time `db:"d2_expire_at"`
	DDonor2     bool      `db:"d2_donor"`
	PNumber2    string    `db:"p2_number"`
	PCountry2   string    `db:"p2_country"`
	PSignature2 []byte    `db:"p2_signature"`
	CreditCards []*user.CreditCard
}
