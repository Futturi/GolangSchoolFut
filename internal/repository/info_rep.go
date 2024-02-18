package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Futturi/GolangSchoolProject/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Info_Repo struct {
	db  *sqlx.DB
	red *redis.Client
}

func NewInfo_Repo(db *sqlx.DB, red *redis.Client) *Info_Repo {
	return &Info_Repo{db: db, red: red}
}

func (r *Info_Repo) Info() (models.Information, error) {
	var namee string
	var abount string
	ctx := context.Background()
	var countU int
	var countL int
	name := r.red.Get(ctx, "Author")
	about := r.red.Get(ctx, "About")
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s", studentTable)
	row := r.db.QueryRow(query)
	if err := row.Scan(&countU); err != nil {
		return models.Information{}, err
	}

	query2 := fmt.Sprintf("SELECT COUNT(id) FROM %s", lessonsTable)
	row2 := r.db.QueryRow(query2)
	if err := row2.Scan(&countL); err != nil {
		return models.Information{}, err
	}
	c, err2 := r.red.Get(ctx, "CountUsers").Result()
	fmt.Println(c)
	c1, err3 := r.red.Get(ctx, "CountLessons").Result()
	if err2 == redis.Nil || c != strconv.Itoa(countU) {
		err := r.red.Set(ctx, "CountUsers", countU, 0)
		if err != nil {
			return models.Information{}, err.Err()
		}
		c = strconv.Itoa(countU)
	}
	if err3 == redis.Nil || c1 != strconv.Itoa(countL) {
		err := r.red.Set(ctx, "CountLessons", countL, 0)
		if err != nil {
			return models.Information{}, err.Err()
		}
		c1 = strconv.Itoa(countL)
	}
	if err := name.Scan(&namee); err != nil {
		return models.Information{}, err
	}
	if err := about.Scan(&abount); err != nil {
		return models.Information{}, err
	}
	result := models.Information{
		Author:       namee,
		About:        abount,
		CountUsers:   c,
		CountLessons: c1,
	}
	return result, nil
}
