package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

//Bank estrutura de dados que se referencia ao banco na nossa aplicacao
type Bank struct {
	Code     string     `json:"code" valid:"notnull"`
	Name     string     `json:"name" valid:"notnull"`
	Accounts []*Account `valid:"-"`
	Base     `valid:"required"`
}

func (b *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(b)
	if err != nil {
		return err
	}
	return nil
}

//NewBank cria uma nova entidade de banco
func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Name: name,
		Code: code,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()
	bank.UpdatedAt = time.Now()

	err := bank.isValid()

	if err != nil {
		return nil, err
	}

	return &bank, nil
}
