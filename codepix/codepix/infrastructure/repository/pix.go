package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/lucasres/code-pix/domain/model"
)

type PixeKeyRepositoryDB struct {
	DB *gorm.DB
}

func (r PixeKeyRepositoryDB) AddBank(bank *model.Bank) error {
	err := r.DB.Create(bank).Error
	if err != nil {
		return err
	}

	return nil
}

func (r PixeKeyRepositoryDB) AddAccount(account *model.Account) error {
	err := r.DB.Create(account).Error
	if err != nil {
		return err
	}

	return nil
}

func (r PixeKeyRepositoryDB) RegisterKey(pixKey *model.PixKey) error {
	err := r.DB.Create(pixKey).Error
	if err != nil {
		return err
	}

	return nil
}

func (r PixeKeyRepositoryDB) FindKeyByKind(kind string, key string) (*model.PixKey, error) {
	var pixKey model.PixKey
	r.DB.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, errors.New("No pixkey find")
	}

	return &pixKey, nil
}

func (r PixeKeyRepositoryDB) FindAccount(id string) (*model.Account, error) {
	var account model.Account
	r.DB.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, errors.New("No pixkey find")
	}

	return &account, nil
}

func (r PixeKeyRepositoryDB) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank
	r.DB.Preload("Bank").First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, errors.New("No pixkey find")
	}

	return &bank, nil
}
