package user

import (
	"rocks/pkg/database/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
    FindAll() ([]*User, error)
}

type repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {
    return &repository{db: db}
}


func (r *repository) FindAll() ([]*User, error) {
    var users []entities.User
    err := r.db.Find(&users).Error
    if err != nil {
        return nil, err
    }

    result := []*User{}
    for _, e := range users {
        result = append(result, &User{
            ID:       e.ID,
            Username: e.Username,
            Email:    e.Email,
            FullName:    e.FullName,
            PhoneNumber:    e.PhoneNumber,
            AvatarURL:    e.AvatarURL,
        })
    }
    return result, nil
}
