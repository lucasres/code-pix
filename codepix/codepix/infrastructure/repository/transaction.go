package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/lucasres/code-pix/domain/model"
)

type TransactionRepositoryDB struct {
	DB *gorm.DB
}

func (r *TransactionRepositoryDB) Register(transaction *model.Transaction) error {
	err := r.DB.Create(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepositoryDB) Save(transaction *model.Transaction) error {
	err := r.DB.Save(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *PixeKeyRepositoryDB) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction
	r.DB.Preload("AccountForm.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, errors.New("Dont find Transaction")
	}

	return &transaction, nil
}
