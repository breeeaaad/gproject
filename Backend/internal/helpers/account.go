package helpers

type Account struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
	Totp     string `json:"totp"`
}

type Totp struct {
	Token string `uri:"token" binding:"required"`
}
