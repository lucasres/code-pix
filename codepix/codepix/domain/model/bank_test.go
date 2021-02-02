package model_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/lucasres/code-pix/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewBank(t *testing.T) {

	code := "001"
	name := "Banco do Brasil"

	//teste que tem que passar
	bank, err := model.NewBank(code, name)
	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(bank.ID))
	require.Equal(t, bank.Code, code)
	require.Equal(t, bank.Name, name)
	
	//teste com falha
	_, err = model.NewBank("", "")
	require.NotNil(t, err)
}