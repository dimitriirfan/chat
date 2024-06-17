package usecase

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/dimitriirfan/chat-2/modules/users/internal/entity"
	"github.com/dimitriirfan/chat-2/modules/users/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const SECRET_KEY = "SECRET_KEY"

type AuthUsecase interface {
	IsTokenValid(ctx context.Context, token string) bool
	Register(ctx context.Context, username string, password string) (entity.AuthRegisterResponse, error)
	Login(ctx context.Context, username string, password string) (entity.AuthLoginResponse, error)
	GetUserFromToken(ctx context.Context, token string) (entity.User, error)
}

type authUc struct {
	usersRepository repository.UsersRepository
}

func NewAuthUsecase(usersRepository repository.UsersRepository) *authUc {
	return &authUc{
		usersRepository: usersRepository,
	}
}

func (uc *authUc) IsTokenValid(ctx context.Context, token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv(SECRET_KEY)), nil
	})

	if err != nil {
		log.Println(err)
	}

	return err == nil

}

func (uc *authUc) Register(ctx context.Context, username string, password string) (entity.AuthRegisterResponse, error) {

	users, err := uc.usersRepository.Search(ctx, entity.SearchUserParams{Username: username})
	if err != nil {
		return entity.AuthRegisterResponse{}, err
	}

	if len(users) > 0 {
		return entity.AuthRegisterResponse{}, nil
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return entity.AuthRegisterResponse{}, err
	}

	newUser := entity.User{
		Username: username,
		Password: string(hashedPass),
	}

	lastInsertId, err := uc.usersRepository.Save(ctx, newUser)
	if err != nil {
		return entity.AuthRegisterResponse{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": lastInsertId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	signedToken, err := token.SignedString([]byte(os.Getenv(SECRET_KEY)))
	if err != nil {
		return entity.AuthRegisterResponse{}, err
	}

	return entity.AuthRegisterResponse{Token: signedToken}, nil
}

func (uc *authUc) Login(ctx context.Context, username string, password string) (entity.AuthLoginResponse, error) {
	users, err := uc.usersRepository.Search(ctx, entity.SearchUserParams{Username: username})
	if err != nil {
		return entity.AuthLoginResponse{}, err
	}

	if len(users) == 0 {
		return entity.AuthLoginResponse{}, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(password)); err != nil {
		return entity.AuthLoginResponse{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": users[0].ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	signedToken, err := token.SignedString([]byte(os.Getenv(SECRET_KEY)))
	if err != nil {
		return entity.AuthLoginResponse{}, err
	}

	return entity.AuthLoginResponse{Token: signedToken}, nil

}

func (uc *authUc) GetUserFromToken(ctx context.Context, token string) (entity.User, error) {

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv(SECRET_KEY)), nil
	})
	if err != nil {
		return entity.User{}, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return entity.User{}, errors.New("invalid token")
	}

	userID := claims["user_id"].(float64)

	return uc.usersRepository.GetByID(ctx, int(userID))
}
