package usecase

import (
	"User-Post-Backend/infra"
	"User-Post-Backend/internal/entity"
	"User-Post-Backend/internal/repository"
	"encoding/json"
	"strconv"
)

type UserUsecase interface {
	Create(user entity.CreateUser) error
	GetAll() ([]entity.User, error)
	GetByID(id uint64) (entity.User, error)
	Update(user entity.User) error
	Delete(id uint64) error
}

type userUsecase struct {
	repo  repository.UserRepository
	cache *infra.RedisClient
}

func NewUserUsecase(repo repository.UserRepository, cache *infra.RedisClient) UserUsecase {
	return &userUsecase{repo: repo, cache: cache}
}

func (u *userUsecase) Create(user entity.CreateUser) error {

	err := u.repo.Create(user)
	if err != nil {
		return err
	}
	return u.cache.Delete("users")
}

func (u *userUsecase) GetAll() ([]entity.User, error) {
	cachedUsers, err := u.cache.Get("users")
	if err == nil && cachedUsers != "" {
		var users []entity.User
		json.Unmarshal([]byte(cachedUsers), &users)
		return users, nil
	}
	users, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}
	cachedData, _ := json.Marshal(users)
	u.cache.Set("users", string(cachedData))
	return users, nil
}

func (u *userUsecase) GetByID(id uint64) (entity.User, error) {
	cachedUser, err := u.cache.Get("user:" + strconv.Itoa(int(id)))
	if err == nil && cachedUser != "" {
		var user entity.User
		json.Unmarshal([]byte(cachedUser), &user)
		return user, nil
	}
	user, err := u.repo.GetByID(id)
	if err != nil {
		return user, err
	}
	cachedData, _ := json.Marshal(user)
	u.cache.Set("user:"+strconv.Itoa(int(id)), string(cachedData))
	return user, nil
}

func (u *userUsecase) Update(user entity.User) error {
	err := u.repo.Update(user)
	if err != nil {
		return err
	}
	u.cache.Delete("user:" + strconv.Itoa(int(user.ID)))
	u.cache.Delete("users")
	return nil
}

func (u *userUsecase) Delete(id uint64) error {
	err := u.repo.Delete(id)
	if err != nil {
		return err
	}
	u.cache.Delete("user:" + strconv.Itoa(int(id)))
	u.cache.Delete("users")
	return nil
}
