package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	// TransactionPedding transacao pendente
	TransactionPedding string = "pending"
	// TransactionCompleted transacao completada
	TransactionCompleted string = "completed"
	// TransactionError transacao com error
	TransactionError string = "error"
	// TransactionConfirmed transacao confirmada
	TransactionConfirmed string = "confirmed"
)

//Transaction representacao de transacoes
type Transaction struct {
	AccountFrom       *Account `json:"account_from" valid:"-"`
	AccountFromID     string   `gorm:"column:account_from_id;type:uuid;" valid:"notnull"`
	Amount            float64  `json:"amount" valid:"notnull"`
	PixKeyTo          *PixKey  `json:"pixkey_to" valid:"-"`
	PixKeyIDTo        string   `gorm:"column:pix_key_id_to;type:uuid;" valid:"notnull"`
	Status            string   `json:"status" valid:"notnull" gorm:"type:varchar(20)"`
	Description       string   `json:"description" valid:"-" gorm:"type:varchar(255)"`
	CancelDescription string   `json:"cancel_description" valid:"-" gorm:"type:varchar(255)"`
	Base              `valid:"required"`
}

// Transactions lista de transacoes
type Transactions struct {
	Transaction []Transaction
}

// TransactionRepositoryInterface contrato para criacao de um repository de transacoes
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

	if t.Amount <= 0 {
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

// Complete funcao para completar a transacao
func (t *Transaction) Complete() error {
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()

	err := t.isValid()
	if err != nil {
		return err
	}

	return nil
}

// Confirm funcao para confirmar a transacao
func (t *Transaction) Confirm() error {
	t.Status = TransactionConfirmed
	t.UpdatedAt = time.Now()

	err := t.isValid()

	return err
}

// Cancel funcao para cancelar a transacao
func (t *Transaction) Cancel(description string) error {
	t.Status = TransactionError
	t.UpdatedAt = time.Now()
	t.CancelDescription = description

	err := t.isValid()

	return err
}

// NewTransaction cria uma nova instancia de transacao
func NewTransaction(account *Account, pixKeyTo *PixKey, amount float64, description string) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom:   account,
		Description:   description,
		PixKeyTo:      pixKeyTo,
		Amount:        amount,
		Status:        TransactionPedding,
		AccountFromID: account.ID,
		PixKeyIDTo:    pixKeyTo.ID,
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
