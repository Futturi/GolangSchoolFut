package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/Futturi/GolangSchoolProject/internal/models"
	"github.com/Futturi/GolangSchoolProject/internal/repository"
	"github.com/golang-jwt/jwt"
)

const (
	salt   = "erogjnrgjjlsa2oqkjpo12j0i13ju491u3hrijwfjnf"
	jwtKey = "fijgnweijndo2ke21piojr0i23hg9usdobijdnsldkpoqif"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (h *AuthService) SignUp(mod models.Teacher) (int, error) {
	mod.Password = hashPass(mod.Password)
	return h.repo.SignUp(mod)
}

func hashPass(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

type Claims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

func (h *AuthService) SignIn(mod models.SignInTeacher) (string, string, error) {
	refresh := h.GenerateRefresh(mod.Username, mod.Password)
	fmt.Println(refresh)
	mod.Password = hashPass(mod.Password)
	id, err := h.repo.SignIn(mod, refresh, time.Now().Add(24*30*time.Hour).Unix())
	if err != nil {
		return "", "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{id, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		IssuedAt:  time.Now().Unix(),
	}})

	jwttoken, err := token.SignedString([]byte(jwtKey))
	return refresh, jwttoken, err
}

func (h *AuthService) ParseToken(header string) (int, error) {
	token, err := jwt.ParseWithClaims(header, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, errors.New("invalid signing method")
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		return 0, err
	}
	Claims, ok := token.Claims.(*Claims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return Claims.Id, nil
}

func (h *AuthService) GenerateRefresh(username, password string) string {
	result := make([]byte, 32)
	rand.Read(result)
	return fmt.Sprintf("%x", result)
}

func (h *AuthService) RefreshToken(refresh string) (string, error) {
	id, err := h.repo.GetByRefresh(refresh)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{id, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		IssuedAt:  time.Now().Unix(),
	}})
	return token.SignedString([]byte(jwtKey))
}
