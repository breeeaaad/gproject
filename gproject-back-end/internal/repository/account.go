package repository

import (
	"github.com/breeeaaad/gproject/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (r *Repository) Check(user models.Account) (int, string, bool, any, error) {
	var (
		id       int
		usern    string
		hash     string
		is_admin bool
		totp     any
	)
	if err := r.conn.QueryRow(r.context, "select id,username,hash,is_admin,tfa from Account where username=$1", user.User).Scan(&id, &usern, &hash, &is_admin, &totp); err != nil {
		return 0, "", false, nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(user.Password)); err != nil {
		return 0, "", false, nil, err
	}
	return id, usern, is_admin, totp, nil
}

func (r *Repository) Add(user models.Account) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 16)
	if err != nil {
		return err
	}
	if _, err := r.conn.Exec(r.context, "insert into Account(username,hash) values ($1,$2)", user.User, string(hash)); err != nil {
		return err
	}
	return nil
}
