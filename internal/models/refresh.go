package models

type Refresh struct {
	Token string `json:"refresh_token" db:"refresh_token"`
}
