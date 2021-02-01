package model

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"time"
	"github.com/satori/go.uuid"
)

//PixKey representacao de uma chave pix no sistema
type PixKey struct {
	Kind string `json:"kind" valid:"notnul"`
	Key string `json:"key" valid:"notnull"`
	AccountID string `json:"account_id" valid:"notnull"`
	Account *Account `json:"account" valid:"-"`
	Status string `json:"status" valid:"notnull"`
	Base `valid:"required"`
}

func (p *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(p)

	if(err != nil){
		return err
	}

	if (p.Kind != "email" && p.Kind != "CPF") {
		return errors.New("Invalid type of Kind")
	}

	if (p.Status != "active" && p.Kind != "inactive") {
		return errors.New("Invalid type of Status")
	}

	return nil
}

func newPixe(kind string, key string, account *Account) (*PixKey, error) {
	pixKey := PixKey{
		Key: key,
		Kind: kind,
		Account: account,
		Status: "active",
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()
	pixKey.UpdatedAt = time.Now()
	pixKey.AccountID = account.ID

	err := pixKey.isValid()
	if(err != nil){
		return nil ,err
	}

	return &pixKey, nil
}