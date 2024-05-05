package services

import (
	"github.com/mateusrlopez/url-shortener-server/models"
	"github.com/mateusrlopez/url-shortener-server/repositories"
	"github.com/mateusrlopez/url-shortener-server/utils"
)

type Records struct {
	repository  *repositories.Records
	idGenerator utils.IdGenerator
	encryptor   utils.Encryptor
}

func NewRecords(repository *repositories.Records, idGenerator utils.IdGenerator, encryptor utils.Encryptor) *Records {
	return &Records{
		repository:  repository,
		idGenerator: idGenerator,
		encryptor:   encryptor,
	}
}

func (s Records) Create(longURL string) (models.Record, error) {
	id := s.idGenerator.GenerateId()
	shortURLKey := s.encryptor.EncryptId(id)

	record, err := s.repository.Insert(id, shortURLKey, longURL)
	if err != nil {
		return models.Record{}, err
	}

	return record, nil
}

func (s Records) FindOneByKey(shortURLKey string) (models.Record, error) {
	record, err := s.repository.FindOneByKey(shortURLKey)
	if err != nil {
		return models.Record{}, err
	}

	return record, nil
}

func (s Records) FindOneAndRegisterAccessByKey(shortURLKey string) (models.Record, error) {
	record, err := s.repository.FindOneAndUpdateAccessesByKey(shortURLKey)
	if err != nil {
		return models.Record{}, err
	}

	return record, nil
}
