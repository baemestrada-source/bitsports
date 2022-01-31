package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email 	 string `json:"email"`
	Name 	 string `json:"name"`
}

type Token struct {
	Username  string `json:"username"`		
	Password  string `json:"password"`		
	Token     string `json:"token"`
	ExpiredAt int    `json:"expired_at"`
}