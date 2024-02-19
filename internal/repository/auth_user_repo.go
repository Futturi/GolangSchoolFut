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
	query := fmt.Sprintf("INSERT INTO %s(username, email, password_hash, token_email) VALUES($1, $2, $3, $4) RETURNING ID", studentTable)
	row := r.db.QueryRow(query, user.Username, user.Email, user.Password, user.Token)
	if err := row.Scan(&id); err != nil {
		return "", nil
	}
	return strconv.Itoa(id), nil
}

func (r *Authorization_User) SignInStudent(userlog models.SignInStudent, refresh string, exp int64) (int, error) {
	var id int
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 and password_hash = $2", studentTable)
	row := tx.QueryRow(query, userlog.Username, userlog.Password)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	query2 := fmt.Sprintf("UPDATE %s SET refresh_token = $1, refresh_token_exxpiry = $2 WHERE id = $3", studentTable)
	_, err = tx.Exec(query2, refresh, exp, id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Authorization_User) CheckHealth(user_id int) int {
	var health int
	query := fmt.Sprintf("SELECT health FROM %s WHERE id = $1", studentTable)
	row := r.db.QueryRow(query, user_id)
	if err := row.Scan(&health); err != nil {
		return 0
	}
	return health
}

func (r *Authorization_User) GetIdByRefresh(refresh models.Refresh) (int, int64, error) {
	var id int
	var time int64
	query := fmt.Sprintf("SELECT id, refresh_token_exxpiry FROM %s WHERE refresh_token = $1", studentTable)
	row := r.db.QueryRow(query, refresh.Token)
	if err := row.Scan(&id, &time); err != nil {
		return 0, 0, err
	}
	return id, time, nil
}

func (r *Authorization_User) CheckToken(token string) error {
	query := fmt.Sprintf("UPDATE %s SET is_email_verified = true WHERE token_email = $1", studentTable)
	_, err := r.db.Exec(query, token)
	if err != nil {
		return err
	}
	return nil
}

func (r *Authorization_User) CheckVer(userlog models.SignInStudent) (bool, error) {
	var ver bool
	query := fmt.Sprintf("SELECT is_email_verified FROM %s WHERE username = $1", studentTable)
	row := r.db.QueryRow(query, userlog.Username)
	if err := row.Scan(&ver); err != nil {
		return false, err
	}
	return ver, nil
}
