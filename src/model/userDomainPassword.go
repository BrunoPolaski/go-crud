package model

import "golang.org/x/crypto/bcrypt"

func (ud *userDomain) EncryptPassword() error {
	if hash, err := bcrypt.GenerateFromPassword([]byte(ud.password), bcrypt.DefaultCost); err != nil {
		return err
	} else {
		ud.password = string(hash)
		return nil
	}
}
