package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	TransactionPedding   string = "pedding"
	TransactionCompleted string = "completed"
	TransactionError     string = "error"
	TransactionConfirmed string = "confirmed"
)

type Transaction struct {
	AccountFrom       *Account `json:"account_from" valid:"notnull"`
	Amount            float64  `json:"amount" valid:"notnull"`
	PixKeyTo          *PixKey  `json:"pixkey_to" valid:"-"`
	Status            string   `json:"status" valid:"notnull"`
	Description       string   `json:"description" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" valid:"-"`
	Base              `valid:"required"`
}

type Transactions struct {
	Transaction []Transaction
}

type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) (*Transaction, error)
	Save(transaction *Transaction) (*Transaction, error)
	Find(id string) (*Transaction, error)
}

func (t *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(t)

	if err != nil {
		return err
	}

	if t.Amount >= 0 {
		return errors.New("Amount must great of than 0")
	}

	if t.Status != TransactionCompleted && t.Status != TransactionError && t.Status != TransactionConfirmed && t.Status != TransactionPedding {
		return errors.New("Invalid type of Status")
	}

	if t.PixKeyTo.AccountID == t.AccountFrom.ID {
		return errors.New("The source and destination account dont be the same")
	}

	return nil
}

func (t *Transaction) Complete() error {
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()

	err := t.isValid()
	if err != nil {
		return err
	}

	return nil
}

func (t *Transaction) Confirm() error {
	t.Status = TransactionConfirmed
	t.UpdatedAt = time.Now()

	err := t.isValid()

	return err
}

func (t *Transaction) Cancel(description string) error {
	t.Status = TransactionError
	t.UpdatedAt = time.Now()
	t.Description = description

	err := t.isValid()

	return err
}

func NewTransaction(account *Account, pixKeyTo *PixKey, amount float64, description string) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom: account,
		Description: description,
		PixKeyTo:    pixKeyTo,
		Amount:      amount,
		Status:      TransactionPedding,
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
