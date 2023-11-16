package repository

func (r *Repository) Authtoken(totp string, user string) error {
	_, err := r.conn.Exec(r.context, "update Account set 2fa=$1 where user=$2 and 2fa is null", totp, user)
	return err
}
