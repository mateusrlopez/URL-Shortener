package clients

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresClient(dsn string) (*gorm.DB, error) {
	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}

	return client, nil
}

func DisconnectPostgresClient(client *gorm.DB) error {
	sql, err := client.DB()

	if err != nil {
		return err
	}

	return sql.Close()
}
