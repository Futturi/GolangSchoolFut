package service

import (
	"errors"
	"fmt"
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

func (a *AuthServiceUser) SignUpStudent(user models.Student, cfg Sender) (string, error) {
	token := utils.GenerateTokenForAccess()
	user.Token = token
	Sendmail(cfg, user.Email, token)
	user = models.Student{Username: user.Username, Email: user.Email, Password: utils.HashPass(user.Password)}
	return a.repo.SignUpStudent(user)
}

func Sendmail(cfg Sender, tomail, token string) {
	sender := NewEmailSender(cfg.EmailSenderName, cfg.EmailSenderAddress, cfg.EmailSenderPassword)

	text := "Authorization to Schoool"
	content := fmt.Sprintf("Authorization link: localhost:8000/auth/%s", token)

	to := []string{tomail}

	err := sender.Sendmail(text, content, to, nil, nil, nil)
	if err != nil {
		fmt.Println("error while sending mail")
	}
}

type ClaimsUser struct {
	Id int
	jwt.StandardClaims
}

func (a *AuthServiceUser) SignInStudent(userlog models.SignInStudent) (string, string, error) {
	refresh := utils.GenerateRefresh()
	userlog = models.SignInStudent{Username: userlog.Username, Password: utils.HashPass(userlog.Password)}
	id, err := a.repo.SignInStudent(userlog, refresh, time.Now().Add(24*30*time.Hour).Unix())
	verified, err := a.repo.CheckVer(userlog)
	if err != nil {
		return "", "", err
	}
	if !verified {
		return "", "", errors.New("your account is not verified")
	}
	if err != nil {
		return "", "", err
	}
	Claims := ClaimsUser{id, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		IssuedAt:  time.Now().Unix(),
	}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	access, err := token.SignedString([]byte(jwtKey1))
	return access, refresh, err
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

func (h *AuthServiceUser) CheckHealth(user_id int) int {
	return h.repo.CheckHealth(user_id)
}

func (h *AuthServiceUser) RefreshUser(refresh models.Refresh) (string, error) {
	id, exp, err := h.repo.GetIdByRefresh(refresh)
	if err != nil {
		return "", err
	}
	if exp < time.Now().Unix() {
		return "", errors.New("your refresh token is expired")
	}
	Claims := ClaimsUser{id, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		IssuedAt:  time.Now().Unix(),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	return token.SignedString([]byte(jwtKey1))
}

func (h *AuthServiceUser) CheckToken(token string) error {
	return h.repo.CheckToken(token)
}
