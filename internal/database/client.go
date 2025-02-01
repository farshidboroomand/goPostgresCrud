package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DatabaseCient interface {
	Ready() bool
}

type Client struct {
	Db *gorm.DB
}

// Ready implements DatabaseCient.
func (c Client) Ready() bool {
	var ready string
	tx := c.Db.Raw("SELECT 1 as ready").Scan(&ready)
	if tx.Error != nil {
		return false
	}

	if ready == "1" {
		return true
	}

	return false
}

func NewDatabaseClient() (DatabaseCient, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		"localhost",
		"5432",
		"postgres",
		"postgres",
		"password",
		"disable",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "wisdom.",
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		QueryFields: true,
	})
	if err != nil {
		return nil, err
	}
	client := Client{
		Db: db,
	}
	return client, nil
}
