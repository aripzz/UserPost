package repository

import (
	"User-Post-Backend/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user entity.CreateUser) error
	GetAll() ([]entity.User, error)
	GetByID(id uint64) (entity.User, error)
	Update(user entity.User) error
	Delete(id uint64) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user entity.CreateUser) error {
	newUser := entity.User{
		Name: user.Name,
	}
	return r.db.Create(&newUser).Error
}

func (r *userRepository) GetAll() ([]entity.User, error) {
	var users []entity.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetByID(id uint64) (entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Update(user entity.User) error {
	return r.db.Model(&user).Where("id = ?", user.ID).Updates(user).Error

}

func (r *userRepository) Delete(id uint64) error {
	if err := r.db.Where("user_id = ?", id).Delete(&entity.Post{}).Error; err != nil {
		return err
	}
	return r.db.Delete(&entity.User{}, id).Error
}
