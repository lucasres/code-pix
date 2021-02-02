package main

import (
	"fmt"
	"log"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID    string `valid:"notnull"`
	Name  string `valid:"notnull"`
	Email string `valid:"notnull"`
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		return err
	}
	return nil
}

func NewUser(name string, email string) (*User, error) {
	user := User{
		Name:  name,
		Email: email,
	}

	u, _ := uuid.NewV4()
	user.ID = u.String()

	err := user.isValid()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func main() {
	user, err := NewUser("Lucas Resende", "lucas@enail.com")
	if err != nil {
		log.Fatalf("Erro em criar um usuario", err)
	}

	fmt.Println("ID    : ", user.ID)
	fmt.Println("Nome  : ", user.Name)
	fmt.Println("Email : ", user.Email)

}
