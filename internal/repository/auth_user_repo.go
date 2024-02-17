package repository

import (
	"fmt"
	"strconv"

	"github.com/Futturi/GolangSchoolProject/internal/models"
	"github.com/jmoiron/sqlx"
)

type Authorization_User struct {
	db *sqlx.DB
}

func NewAuthorization_User(db *sqlx.DB) *Authorization_User {
	return &Authorization_User{db: db}
}

func (r *Authorization_User) SignUpStudent(user models.Student) (string, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s(username, email, password_hash) VALUES($1, $2, $3) RETURNING ID", studentTable)
	row := r.db.QueryRow(query, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return "", nil
	}
	return strconv.Itoa(id), nil
}

func (r *Authorization_User) SignInStudent(userlog models.SignInStudent) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 and password_hash = $2", studentTable)
	row := r.db.QueryRow(query, userlog.Username, userlog.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
