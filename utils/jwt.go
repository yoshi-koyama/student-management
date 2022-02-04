package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

var SecretKey = ""
var STUDENT_ID = "student_id"

func CreateToken(studentId int64) (string, error) {
	claim := jwt.MapClaims{}
	claim[STUDENT_ID] = studentId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", fmt.Errorf("cannot generate token")
	}

	return signedToken, nil
}

func ValidateToken(value string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(value, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claim, nil
	} else {
		return nil, err
	}

}
