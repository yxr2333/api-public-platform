package security

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(username string, id uint, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
	issure    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: "secret",
		issure:    "test.com",
	}
}

func (service *jwtService) GenerateToken(username string, id uint, isUser bool) string {
	claims := jwt.MapClaims{}
	claims["authorized"] = isUser
	claims["username"] = username
	claims["userId"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	claims["iss"] = service.issure
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	encodedToken, _ := token.SignedString([]byte(service.secretKey))
	return encodedToken
}

func (service *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}
