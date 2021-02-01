package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

//Account representacao de uma conta de um banco no sistema
type Account struct {
	OwnerName string `json:"owner_name" valid:"notnull"`
	Bank *Bank `json:"bank" valid:"-"`
	Number string `json:"number" valid:"notnull"`
	Base `valid:"required"`
} 

func (a *Account) isValid() error {
	_, err := govalidator.ValidateStruct(a)
	if(err != nil) {
		return err
	}
	return nil
}

//NewAccount cria a entidade uma conta
func NewAccount(bank *Bank, number string, ownerName string) (*Account, error) {
	account := Account{
		Bank: bank,
		Number: number,
		OwnerName: ownerName,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()

	err := account.isValid()
	if(err != nil){
		return nil, err
	}

	return &account, nil
}