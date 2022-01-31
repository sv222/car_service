package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int64  `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

func (u *User) EncryptPassword() error {
	if len(u.Password) > 8 {
		encPass, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = encPass
	}

	return nil
}

func encryptString(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
