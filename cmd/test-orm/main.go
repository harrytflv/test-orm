package main

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/harrytflv/test-orm/pkg/user"
	"github.com/harrytflv/test-orm/pkg/utils"
	"github.com/upper/db/v4/adapter/cockroachdb"
)

var (
	dsn = "host=192.168.50.45 user=test dbname=test port=26257 sslmode=disable"

	settings = cockroachdb.ConnectionURL{
		Host:     "192.168.50.45",
		Database: "test",
		User:     "test",
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
		CreditCards: []*user.CreditCard{discover, chase},
	}

	userNum   = 10000
	workerNum = 10
)

func main() {
	upperDbOrm()
	upperDbSqlBuilder()
	gormOrm()
}

func gormOrm() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("cockroachdb.Open: ", err)
	}
	log.Println("cockroachdb connected")
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("cockroachdb.Open: ", err)
	}
	defer sqlDB.Close()

	start := time.Now()
	utils.Work(userNum, workerNum, func() {
		tmpJack := *jack
		tmpJack.ID = utils.UUID()
		db.Create(jack)
	})
	end := time.Now()
	log.Printf("gormOrm:           %v\n", end.Sub(start))
}

func upperDbOrm() {
	sess, err := cockroachdb.Open(settings)
	if err != nil {
		log.Fatal("cockroachdb.Open: ", err)
	}
	log.Println("cockroachdb connected")
	defer sess.Close()

	start := time.Now()
	utils.Work(userNum, workerNum, func() {
		tmpJack := *jack
		tmpJack.ID = utils.UUID()
		_, err := sess.Collection("users").Insert(tmpJack)
		if err != nil {
			log.Fatal("sess.Save: ", err)
		}
	})
	end := time.Now()
	log.Printf("upperDbOrm:        %v\n", end.Sub(start))
}

func upperDbSqlBuilder() {
	sess, err := cockroachdb.Open(settings)
	if err != nil {
		log.Fatal("cockroachdb.Open: ", err)
	}
	log.Println("cockroachdb connected")
	defer sess.Close()

	start := time.Now()
	utils.Work(userNum, workerNum, func() {
		tmpJack := *jack
		tmpJack.ID = utils.UUID()
		if _, err := sess.SQL().InsertInto("users").Columns(
			"id",
			"created_at",
			"updated_at",
			"last_name",
			"first_name",
			"nick_name",
			"date_of_birth",
			"drivers_license",
			"passport",
		).Values(
			tmpJack.ID,
			tmpJack.CreatedAt,
			tmpJack.UpdatedAt,
			tmpJack.LastName,
			tmpJack.FirstName,
			tmpJack.NickName,
			tmpJack.DateOfBirth,
			tmpJack.DriversLicense,
			tmpJack.Passport,
		).Exec(); err != nil {
			log.Fatal("sess.Save: ", err)
		}
	})
	end := time.Now()

	log.Printf("upperDbSqlBuilder: %v\n", end.Sub(start))
}
