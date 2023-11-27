package repository

import (
	"errors"
	"time"

	"github.com/breeeaaad/gproject/internal/configs"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (r *Repository) Genjwt(id int, user string, is_admin bool) (string, string, error) {
	prvKey, err := configs.JwtPrvKey()
	if err != nil {
		return "", "", err
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(prvKey)
	if err != nil {
		return "", "", err
	}
	claims := make(jwt.MapClaims)
	claims["id"] = id
	claims["user"] = user
	claims["is_admin"] = is_admin
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	access, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", "", err
	}
	var refresh string
	if err := r.conn.QueryRow(r.context, "insert into Session(user_id) values($1)", id).Scan(&refresh); err != nil {
		return "", "", err
	}
	return access, refresh, nil
}

func (r *Repository) DelRefresh(refresh string) (bool, error) {
	if tag, err := r.conn.Exec(r.context, "delete from Session where refresh=$1 and expiresIn>now()", refresh); err != nil {
		return false, err
	} else if tag.RowsAffected() != 1 {
		return false, errors.New("No refresh token")
	}
	return true, nil
}

func (r *Repository) Refresh(refresh string) (int, string, bool, error) {
	uuid, err := uuid.Parse(refresh)
	if err != nil {
		return 0, "", false, err
	}
	var (
		id       int
		user     string
		is_admin bool
	)
	if err := r.conn.QueryRow(r.context, "select Account.id,user,is_admin from Account join Session on Account.id=Session.user_id where refresh=$1", uuid).Scan(&id, &user, &is_admin); err != nil {
		return 0, "", false, err
	}
	return id, user, is_admin, nil
}
