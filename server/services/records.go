package services

import (
	"github.com/mateusrlopez/url-shortener-server/models"
	"github.com/mateusrlopez/url-shortener-server/repositories"
)

type Records struct {
	repository *repositories.Records
}

func NewRecords(repository *repositories.Records) *Records {
	return &Records{
		repository: repository,
	}
}

func (s Records) Create(url string) (string, error) {
	key, err := s.repository.Insert(url)
	if err != nil {
		return "", err
	}

	return key, nil
}

func (s Records) FindOneByKey(key string) (models.Record, error) {
	record, err := s.repository.FindOneByKey(key)
	if err != nil {
		return models.Record{}, err
	}

	return record, nil
}

func (s Records) FindOneAndRegisterAccessByKey(key string) (string, error) {
	record, err := s.repository.FindOneAndUpdateAccessesByKey(key)
	if err != nil {
		return "", err
	}

	return record.URL, nil
}
