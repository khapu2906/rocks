package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
    ErrUserExists         = errors.New("username already exists")
    ErrInvalidCredentials = errors.New("invalid credentials")
)

type Usecase interface {
    Get() ([]*User, error)
}

type usecase struct {
    repo UserRepository
}

func NewUsecase(r UserRepository) Usecase {
    return &usecase{r}
}


func (u *usecase) Get() ([]*User, error) {
    uArr, err := u.repo.FindAll()

    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrInvalidCredentials
        }
        return nil, err
    }

    return uArr, nil
}
