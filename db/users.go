package db

import (
	"encoding/base64"
	"fmt"

	"github.com/gophergala2016/wwcdc_01/models"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(user *models.UserRegistration) error {
	query := "insert into gophergala_users (name, email, passhash) values($1, $2, $3)"
	passhash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	base64hash := base64.StdEncoding.EncodeToString(passhash)
	_, err = db.Exec(query, user.Name, user.Email, base64hash)
	if err != nil {
		return err
	}
	return nil
}

func AuthUser(email string, password string) bool {
	query := "select passhash from gophergala_users where email = $1"
	base64hash := ""
	err := db.Get(&base64hash, query, email)
	if err != nil {
		fmt.Println(err)
		return false
	}
	passhash, err := base64.StdEncoding.DecodeString(base64hash)
	if err != nil {
		fmt.Println(err)
		return false
	}
	err = bcrypt.CompareHashAndPassword(passhash, []byte(password))
	if err == nil {
		return true
	}
	return false
}
