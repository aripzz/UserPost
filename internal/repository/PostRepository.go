package repository

import (
	"User-Post-Backend/internal/entity"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post entity.CreatePost) error
	GetAll() ([]entity.Post, error)
	GetByID(id uint64) (entity.Post, error)
	Update(post entity.Post) error
	Delete(id uint64) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(post entity.CreatePost) error {
	newPost := entity.Post{
		Title:   post.Title,
		Content: post.Content,
		UserID:  post.UserID,
	}
	return r.db.Create(&newPost).Error
}

func (r *postRepository) GetAll() ([]entity.Post, error) {
	var posts []entity.Post
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) GetByID(id uint64) (entity.Post, error) {
	var post entity.Post
	if err := r.db.First(&post, id).Error; err != nil {
		return post, err
	}
	return post, nil
}

func (r *postRepository) Update(post entity.Post) error {
	return r.db.Save(&post).Error
}

func (r *postRepository) Delete(id uint64) error {
	return r.db.Delete(&entity.Post{}, id).Error
}
