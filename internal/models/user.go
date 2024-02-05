package models

type Teacher struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password_hash"`
	Email    string `json:"email" db:"email"`
}

type Student struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password_hash"`
	Email    string `json:"email" db:"email"`
}
type SignInTeacher struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password_hash"`
}
