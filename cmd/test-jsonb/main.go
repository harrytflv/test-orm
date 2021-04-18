package main

import (
	"log"
	"time"

	"github.com/upper/db/v4/adapter/cockroachdb"

	"github.com/harrytflv/test-orm/pkg/user"
	"github.com/harrytflv/test-orm/pkg/userinline"
	"github.com/harrytflv/test-orm/pkg/utils"
)

var (
	settings = cockroachdb.ConnectionURL{
		Host:     "localhost",
		Database: "defaultdb",
		User:     "root",
		Options: map[string]string{
			"sslmode": "disable",
		},
	}

	discover = &user.CreditCard{
		CardNumber:   "1000",
		SecurityCode: 7,
	}

	chase = &user.CreditCard{
		CardNumber:   "2000",
		SecurityCode: 8,
	}

	jack = &user.User{
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		LastName:    "Chen",
		FirstName:   "Jack",
		NickName:    "Piggy",
		DateOfBirth: "10/09/1995",
		DriversLicense: &user.DriversLicense{
			ID:       "F077",
			Class:    "C",
			ExpireAt: time.Now().Add(10 * time.Second),
			Donor:    false,
		},
		Passport: &user.Passport{
			Number:    "F077",
			Country:   "CA",
			Signature: []byte("foo"),
		},
		DriversLicense2: &user.DriversLicense{
			ID:       "F077",
			Class:    "C",
			ExpireAt: time.Now().Add(10 * time.Second),
			Donor:    false,
		},
		Passport2: &user.Passport{
			Number:    "F077",
			Country:   "CA",
			Signature: []byte("foo"),
		},
		CreditCards: []*user.CreditCard{discover, chase},
	}

	jackinline = &userinline.User{
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		LastName:    "Chen",
		FirstName:   "Jack",
		NickName:    "Piggy",
		DateOfBirth: "10/09/1995",
		DID:         "F077",
		DClass:      "C",
		DExpireAt:   time.Now().Add(10 * time.Second),
		DDonor:      false,
		PNumber:     "F077",
		PCountry:    "CA",
		PSignature:  []byte("foo"),
		DID2:        "F077",
		DClass2:     "C",
		DExpireAt2:  time.Now().Add(10 * time.Second),
		DDonor2:     false,
		PNumber2:    "F077",
		PCountry2:   "CA",
		PSignature2: []byte("foo"),
		CreditCards: []*user.CreditCard{discover, chase},
	}

	userNum   = 10000
	listerNum = 10
	workerNum = 8
)

func main() {
	jsonb()
	// inline()
}

func jsonb() {
	sess, err := cockroachdb.Open(settings)
	if err != nil {
		log.Fatal("cockroachdb.Open: ", err)
	}
	log.Println("cockroachdb connected")
	defer sess.Close()

	took := utils.Work(userNum, workerNum, func() {
		_, err := sess.Collection("users").Insert(jack)
		if err != nil {
			log.Fatal("sess.Save: ", err)
		}
	})
	log.Printf("jsonb create:        %v\n", took)

	took = utils.Work(listerNum, workerNum, func() {
		var users []*user.User
		err := sess.Collection("users").Find().All(&users)
		if err != nil {
			log.Fatal("sess.Find: ", err)
		}
	})
	log.Printf("jsonb list:        %v\n", took)
}

func inline() {
	sess, err := cockroachdb.Open(settings)
	if err != nil {
		log.Fatal("cockroachdb.Open: ", err)
	}
	log.Println("cockroachdb connected")
	defer sess.Close()

	took := utils.Work(userNum, workerNum, func() {
		_, err := sess.Collection("usersinline").Insert(jackinline)
		if err != nil {
			log.Fatal("sess.Save: ", err)
		}
	})
	log.Printf("inline create:        %v\n", took)

	took = utils.Work(listerNum, workerNum, func() {
		var users []*userinline.User
		err := sess.Collection("usersinline").Find().All(&users)
		if err != nil {
			log.Fatal("sess.Find: ", err)
		}
	})
	log.Printf("inline list:        %v\n", took)
}
