package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/Futturi/GolangSchoolProject/internal/models"
	"github.com/jmoiron/sqlx"
)

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) SignUp(mod models.Teacher) (int, error) {
	var result int
	query := r.db.QueryRow(fmt.Sprintf(`
	INSERT INTO %s(username, password_hash, email) 
	VALUES($1, $2, $3) RETURNING id`, teachersTable), mod.Username, mod.Password, mod.Email)
	err := query.Scan(&result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (r *AuthRepo) SignIn(mod models.SignInTeacher, refresh string, timerefresh int64) (int, error) {
	var id int
	query_refresh := fmt.Sprintf(`UPDATE %s SET
	refresh_token = $1, refresh_token_exxpiry = $2`, teachersTable)
	_, err := r.db.Exec(query_refresh, refresh, timerefresh)
	if err != nil {
		return 0, err
	}
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND password_hash = $2", teachersTable)
	row := r.db.QueryRow(query, mod.Username, mod.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthRepo) GetByRefresh(refresh string) (int, error) {
	var id int
	var timerefresh int64
	query := fmt.Sprintf("SELECT id, refresh_token_exxpiry FROM %s WHERE refresh_token = $1", teachersTable)
	row := r.db.QueryRow(query, refresh)
	if err := row.Scan(&id, &timerefresh); err != nil {
		return 0, err
	}

	fmt.Println(timerefresh)

	if timerefresh < time.Now().Unix() {
		return 0, errors.New("your refresh token is expired")
	}

	return id, nil
}
