package service

import (
	"time"
	"user-service/src/model"
	"user-service/src/repository"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func CreateToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !tkn.Valid {
		return nil, err
	}

	return claims, nil
}

func SignupUser(username, password string) error {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	user := model.User{Username: username, Password: hashedPassword}
	return repository.CreateUser(&user).Error
}

func LoginUser(username, password string) (string, error) {
	var user model.User
	err := repository.GetUserByUsername(username, &user).Error
	if err != nil {
		return "", err
	}

	err = CheckPasswordHash(password, user.Password)
	if err != nil {
		return "", err
	}

	return CreateToken(username)
}

func UpdateUser(id, username, password string) error {
	var user model.User
	err := repository.GetUserByID(id, &user).Error
	if err != nil {
		return err
	}

	if username != "" {
		user.Username = username
	}

	if password != "" {
		hashedPassword, err := HashPassword(password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}

	return repository.UpdateUser(&user).Error
}

func DeleteUser(id string) error {
	var user model.User
	err := repository.GetUserByID(id, &user).Error
	if err != nil {
		return err
	}

	return repository.DeleteUser(&user).Error
}