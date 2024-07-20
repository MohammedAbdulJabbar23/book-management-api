package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/MohammedAbdulJabbar23/book-management-api/config"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)



var jwtKey = []byte("HELLOAcatWhoATEIT'STAILBUTHAPPY");


type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func (user *User) Register() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost);
	if err != nil {
		return err;
	}
	user.Password = string(hashedPassword);
	_,err = config.DB.Exec(`INSERT INTO users(username,password) VALUES ($1,$2)`, user.Username, user.Password);
	
	return err;
}

func (user *User) Authenticate() (string, error) {
	var dbUser User;
	err := config.DB.QueryRow("SELECT id, username, password FROM users WHERE username=$1", user.Username).Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password);
	if err != nil {
		fmt.Println("#1");
		if err == sql.ErrNoRows {
			return "", errors.New("invalid credentials");
		}
		return "",err;
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password),[]byte(user.Password)); err != nil {
		fmt.Println("#2");
		return "", errors.New("invalid credentials");
	}
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID: dbUser.ID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

	return tokenString, nil;
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{};
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token)(interface{}, error) {
		return jwtKey, nil;
	});
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token");
	}
	return claims, nil;
}