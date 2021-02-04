package usecase

import "github.com/lucasres/code-pix/domain/model"

//TransactionUseCase Implementacao do caso de uso
type TransactionUseCase struct {
	Repository    model.TransactionRepositoryInterface
	PixRepository model.PixKeyRepositoryInterface
}

//Register cria uma nova transaction
func (t *TransactionUseCase) Register(accountID string, amount float64, pixKeyTo string, pixKeyKind string, description string) (*model.Transaction, error) {
	account, err := t.PixRepository.FindAccount(accountID)
	if err != nil {
		return nil, err
	}

	pixKey, err := t.PixRepository.FindKeyByKind(pixKeyTo, pixKeyKind)
	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, pixKey, amount, description)
	if err != nil {
		return nil, err
	}

	transaction, err = t.Repository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

//Confirm confirma uma transaction
func (t *TransactionUseCase) Confirm(transactionID string) (*model.Transaction, error) {
	transaction, err := t.Repository.Find(transactionID)
	if err != nil {
		return nil, err
	}

	transaction.Confirm()
	transaction, err = t.Repository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

//Error uma transaction com error
func (t *TransactionUseCase) Error(transactionID string, reason string) (*model.Transaction, error) {
	transaction, err := t.Repository.Find(transactionID)
	if err != nil {
		return nil, err
	}

	transaction.Cancel(reason)
	transaction, err = t.Repository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
