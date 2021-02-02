package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	PixKeyActiveStatus   string = "active"
	PixKeyInactiveStatus string = "inactive"
	PixKeyKindCPF        string = "cpf"
	PixKeyKindEmail      string = "email"
)

//PixKey representacao de uma chave pix no sistema
type PixKey struct {
	Kind      string   `json:"kind" valid:"notnul"`
	Key       string   `json:"key" valid:"notnull"`
	AccountID string   `json:"account_id" valid:"notnull"`
	Account   *Account `json:"account" valid:"-"`
	Status    string   `json:"status" valid:"notnull"`
	Base      `valid:"required"`
}

type PixKeyRepositoryInterface interface {
	Register(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) (*Bank, error)
	FindAccount(id string) (*Account, error)
	AddAccount(account *Account) (*Account, error)
}

func (p *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return err
	}

	if p.Kind != PixKeyKindCPF && p.Kind != PixKeyKindEmail {
		return errors.New("Invalid type of Kind")
	}

	if p.Status != PixKeyActiveStatus && p.Kind != PixKeyInactiveStatus {
		return errors.New("Invalid type of Status")
	}

	return nil
}

func newPixe(kind string, key string, account *Account) (*PixKey, error) {
	pixKey := PixKey{
		Key:     key,
		Kind:    kind,
		Account: account,
		Status:  PixKeyActiveStatus,
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()
	pixKey.UpdatedAt = time.Now()
	pixKey.AccountID = account.ID

	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}

	return &pixKey, nil
}
