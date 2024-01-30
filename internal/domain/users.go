package domain

type Users struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type UserLogin struct {
	Login string `json:"login"`
}
