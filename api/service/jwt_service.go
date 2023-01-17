package service

// import (
// 	"errors"
// 	"os"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/google/uuid"
// )

// type JWTService interface {
// 	// GenerateToken(userId uuid.UUID, Username string, Email string, RoleId uuid.UUID, Minute int, mongo bool) (string, error)
// 	ValidateAuthToken(token string) (*jwt.Token, error)
// 	// ValidateRefreshToken(token string) (*jwt.Token, error)
// 	GetUserData(token string, data string) (string, error)
// 	// FindByRefreshToken(refreshToken string) (string, error)
// 	// DeleteRefreshToken(refreshToken string)
// }

// type jwtCustomClaims struct {
// 	UserId   uuid.UUID `json:"user_id"`
// 	Username string    `json:"username"`
// 	Email    string    `json:"email"`
// 	RoleId   uuid.UUID `json:"role_id"`

// 	jwt.StandardClaims
// }

// type jwtService struct {
// 	// jwtRepository repository.JwtRepository
// 	secretKeyAuth    string
// 	secretKeyRefresh string
// 	issuer           string
// }

// func NewJWTService() JWTService {
// 	authKey := os.Getenv("SECRET_KEY_AUTH")
// 	refreshKey := os.Getenv("SECRET_KEY_REFRESH")
// 	return &jwtService{
// 		// jwtRepository:    jwtRepository,
// 		issuer:           "golang-jwt",
// 		secretKeyAuth:    authKey,
// 		secretKeyRefresh: refreshKey,
// 	}
// }

// // func (jwtService *jwtService) GenerateToken(UserId uuid.UUID, Username string, Email string, RoleId uuid.UUID, Minute int, inputMongoDB bool) (string, error) {
// // 	claims := &jwtCustomClaims{
// // 		UserId,
// // 		Username,
// // 		Email,
// // 		RoleId,
// // 		jwt.StandardClaims{
// // 			ExpiresAt: time.Now().Add(time.Minute * time.Duration(Minute)).Unix(),
// // 			Issuer:    jwtService.issuer,
// // 			IssuedAt:  time.Now().Unix(),
// // 		},
// // 	}
// // 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// // 	t, err := token.SignedString([]byte(jwtService.secretKeyAuth))
// // 	if err != nil {
// // 		return "", err
// // 	}
// // 	if inputMongoDB {
// // 		err = jwtService.jwtRepository.CreateRefreshToken(t)
// // 		if err != nil {
// // 			return "", err
// // 		}
// // 	}
// // 	return t, nil
// // }

// func (jwtService *jwtService) ValidateAuthToken(token string) (*jwt.Token, error) {
// 	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errors.New("Auth Token is invalid")
// 		}
// 		return []byte(jwtService.secretKeyAuth), nil
// 	})
// }

// // func (jwtService *jwtService) ValidateRefreshToken(token string) (*jwt.Token, error) {
// // 	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
// // 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// // 			return nil, errors.New("Refresh Token is invalid")
// // 		}
// // 		return []byte(jwtService.secretKeyRefresh), nil
// // 	})
// // }

// func (jwtService *jwtService) GetUserData(token string, data string) (string, error) {
// 	jwtToken, err := jwtService.ValidateAuthToken(token)
// 	if err != nil {
// 		return "", err
// 	}
// 	claims := jwtToken.Claims.(jwt.MapClaims)
// 	userData := claims[data].(string)
// 	if err != nil {
// 		return "", err
// 	}
// 	return userData, nil
// }

// // func (jwtService *jwtService) FindByRefreshToken(refreshToken string) (string, error) {
// // 	token, err := jwtService.jwtRepository.FindByRefreshToken(refreshToken)
// // 	if err != nil {
// // 		return "", errors.New("Refresh Token is not found")
// // 	}
// // 	return token, nil
// // }

// // func (jwtService *jwtService) DeleteRefreshToken(refreshToken string) {
// // 	jwtService.jwtRepository.DeleteRefreshToken(refreshToken)
// // }