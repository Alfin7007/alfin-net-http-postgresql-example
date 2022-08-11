package middlewares

import (
	"errors"
	"http/example/config"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID int, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = userID
	claims["email"] = email
	claims["exp"] = time.Now().Add(5 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Secret_JWT()))
}

func JWTTokenCheck(header string) (int, string, error) {
	jwtToken, errExtract := extractBearerToken(header)
	if errExtract != nil {
		return 0, "", errExtract
	}

	token, errParse := parseToken(jwtToken)
	if errParse != nil {
		return 0, "", errParse
	}

	claims, OK := token.Claims.(jwt.MapClaims)
	if OK {
		userID := claims["userID"].(float64)
		email := claims["email"].(string)

		return int(userID), email, nil
	}
	return 0, "", errors.New("failed extract token")
}

func parseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return []byte(config.Secret_JWT()), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}
	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}
	return jwtToken[1], nil
}
