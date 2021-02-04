package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	//PixKeyActiveStatus chave ativa
	PixKeyActiveStatus string = "active"
	//PixKeyInactiveStatus chave inativa
	PixKeyInactiveStatus string = "inactive"
	//PixKeyKindCPF tipo de chave CPF
	PixKeyKindCPF string = "cpf"
	//PixKeyKindEmail tipo de chave email
	PixKeyKindEmail string = "email"
)

//PixKey representacao de uma chave pix no sistema
type PixKey struct {
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	AccountID string   `gorm:"column:account_id;type:uuid;not null" valid:"-"`
	Account   *Account `json:"account" valid:"-"`
	Status    string   `json:"status" valid:"notnull"`
	Base      `valid:"required"`
}

//PixKeyRepositoryInterface contrato de repositorio
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

func NewPixKey(kind string, key string, account *Account) (*PixKey, error) {
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
