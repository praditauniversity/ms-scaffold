package service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateToken(userId string, Username string, Email string, RoleId uint, Minute int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	GetUserData(token string, data string) (string, error)
}

type jwtCustomClaims struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleId   uint   `json:"role_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "golang-jwt",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "md5(rahasia)"
	}
	return secretKey
}

func (jwtService *jwtService) GenerateToken(UserId string, Username string, Email string, RoleId uint, Minute int) (string, error) {
	claims := &jwtCustomClaims{
		UserId,
		Username,
		Email,
		RoleId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(Minute)).Unix(),
			Issuer:    jwtService.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(jwtService.secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (jwtService *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtService.secretKey), nil
	})
}

func (jwtService *jwtService) GetUserData(token string, data string) (string, error) {
	jwtToken, err := jwtService.ValidateToken(token)
	if err != nil {
		return "", err
	}
	claims := jwtToken.Claims.(jwt.MapClaims)
	userDataString := claims[data].(string)
	if err != nil {
		return "", err
	}
	return userDataString, nil
}