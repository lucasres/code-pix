package model

import (
	"time"
	"github.com/asaskevich/govalidator"
)

func init(){
	govalidator.SetFieldsRequiredByDefault(true)
}

//Base estrutura basica de uma entidade
type Base struct {
	ID string `json:"id"`
	CreatedAt time.Time `json:"created_at" valid:"uuid"`
	UpdatedAt time.Time `json:"updated_at"`
}