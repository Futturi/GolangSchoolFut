package service

import (
	"errors"
	"time"

	"github.com/Futturi/GolangSchoolProject/internal/models"
	"github.com/Futturi/GolangSchoolProject/internal/repository"
	"github.com/Futturi/GolangSchoolProject/internal/utils"
	"github.com/golang-jwt/jwt"
)

const (
	jwtKey1 = "oerjgnwojenflkasfgolekrhokqwjedlkmsdlkbnerlkgmqw;"
)

type AuthServiceUser struct {
	repo repository.AuthorizationUser
}

func NewAuthServiceUser(repo repository.AuthorizationUser) *AuthServiceUser {
	return &AuthServiceUser{repo: repo}
}

func (a *AuthServiceUser) SignUpStudent(user models.Student) (string, error) {
	user = models.Student{Username: user.Username, Email: user.Email, Password: utils.HashPass(user.Password)}
	return a.repo.SignUpStudent(user)
}

type ClaimsUser struct {
	Id int
	jwt.StandardClaims
}

func (a *AuthServiceUser) SignInStudent(userlog models.SignInStudent) (string, error) {
	userlog = models.SignInStudent{Username: userlog.Username, Password: utils.HashPass(userlog.Password)}
	id, err := a.repo.SignInStudent(userlog)
	if err != nil {
		return "", err
	}
	Claims := ClaimsUser{id, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		IssuedAt:  time.Now().Unix(),
	}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)

	return token.SignedString([]byte(jwtKey1))
}

func (h *AuthServiceUser) ParseTokenUser(header string) (int, error) {
	token, err := jwt.ParseWithClaims(header, &ClaimsUser{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, errors.New("invalid signing method")
		}
		return []byte(jwtKey1), nil
	})
	if err != nil {
		return 0, err
	}
	Claims, ok := token.Claims.(*ClaimsUser)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return Claims.Id, nil
}
