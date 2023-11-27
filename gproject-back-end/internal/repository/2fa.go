package repository

func (r *Repository) Authtoken(totp string, user string) error {
	_, err := r.conn.Exec(r.context, "update Account set tfa=$1 where username=$2 and tfa is null", totp, user)
	return err
}

func (r *Repository) Deltoken(totp string, user string) error {
	_, err := r.conn.Exec(r.context, "update Account set tfa=null where username=$1 and tfa=$2", user, totp)
	return err
}
