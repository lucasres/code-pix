package usecase

import "github.com/lucasres/code-pix/domain/model"

type PixUseCase struct {
	Repository model.PixKeyRepositoryInterface
}

func (pixUseCase *PixUseCase) RegisterKey(key string, kind string, accountID string) (*model.PixKey, error) {
	account, err := pixUseCase.Repository.FindAccount(accountID)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, key, account)
	if err != nil {
		return nil, err
	}

	pixUseCase.Repository.Register(pixKey)
	if pixKey.ID == "" {
		return nil, err
	}

	return pixKey, nil
}

func (pixUseCase *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := pixUseCase.Repository.FindKeyByKind(key, kind)

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
