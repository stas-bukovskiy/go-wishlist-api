package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/entity"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/repository"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/validation"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/errs"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/logger"
	"time"
)

const (
	salt       = "jevBH89BC9cbdsc298dUCXbzasxOZox"
	singingKey = "pISc0ODSDC9023onc90132sdcu19DUISuisdUSDcx"
	tokenTTL   = 24 * time.Hour
)

type AuthService struct {
	repo repository.User
	log  logger.Logger
}

func NewAuthService(repo repository.User, log logger.Logger) *AuthService {
	return &AuthService{repo: repo, log: log}
}

func (s *AuthService) CreateUser(name, email, password string) (entity.User, error) {
	if !validation.IsValidEmail(email) {
		return entity.User{}, errs.NewError(errs.Validation, "email is not valid")
	}
	user := entity.User{Name: name, Email: email, Password: generatePasswordHash(password)}
	return s.repo.SaveUser(user)
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId uuid.UUID
}

func (s *AuthService) Authenticate(email string, password string) (string, error) {
	user, err := s.repo.GetUserByEmailAndPassword(email, generatePasswordHash(password))
	if err != nil {
		if errs.KindIs(errs.NotFound, err) {
			return "", errs.NewError(errs.Unauthorized, "Invalid email or password")
		}
		return "", err
	}

	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(tokenTTL).Unix(),
			IssuedAt:  now.Unix(),
			Subject:   user.Email,
		},
		UserId: user.ID,
	})

	return token.SignedString([]byte(singingKey))
}

func (s *AuthService) ParseToken(tokenToParse string) (entity.User, error) {
	token, err := jwt.ParseWithClaims(tokenToParse, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.NewError(errs.Unauthorized, "Invalid signing method")
		}

		return []byte(singingKey), nil
	})
	if err != nil {
		return entity.User{}, errs.NewError(errs.Unauthorized, "Invalid access token")
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return entity.User{}, errs.NewError(errs.Unauthorized, "Invalid access token")
	}
	userId := claims.UserId
	user, err := s.repo.GetUserById(userId)
	if err != nil {
		return entity.User{}, errs.NewError(errs.Unauthorized, "Invalid access token")
	}
	return user, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
