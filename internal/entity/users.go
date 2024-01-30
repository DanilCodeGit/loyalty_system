package entity

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Auth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
