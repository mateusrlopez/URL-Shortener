package repositories

import (
	"context"
	"github.com/mateusrlopez/url-shortener-server/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Records struct {
	database *gorm.DB
}

func NewRecords(database *gorm.DB) *Records {
	return &Records{
		database: database,
	}
}

func (r Records) Insert(id int64, shortURLKey, longURL string) (models.Record, error) {
	record := models.Record{
		ID:          id,
		ShortURLKey: shortURLKey,
		LongURL:     longURL,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := r.database.WithContext(ctx).Create(&record).Error; err != nil {
		return models.Record{}, err
	}

	return record, nil
}

func (r Records) FindOneByKey(shortURLKey string) (models.Record, error) {
	var record models.Record

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := r.database.WithContext(ctx).First(&record, "short_url_key = ?", shortURLKey).Error; err != nil {
		return models.Record{}, err
	}

	return record, nil
}

func (r Records) FindOneAndUpdateAccessesByKey(shortURLKey string) (models.Record, error) {
	var record models.Record

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := r.database.WithContext(ctx).Model(&record).Clauses(clause.Returning{}).Where("short_url_key = ?", shortURLKey).Updates(map[string]interface{}{"last_access": time.Now(), "accesses": gorm.Expr("accesses + ?", 1)}).Error; err != nil {
		return models.Record{}, err
	}

	return record, nil
}
