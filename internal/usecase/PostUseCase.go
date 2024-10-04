package usecase

import (
	"User-Post-Backend/infra"
	"User-Post-Backend/internal/entity"
	"User-Post-Backend/internal/repository"
	"encoding/json"
	"errors"
	"strconv"
)

type PostUsecase interface {
	Create(post entity.CreatePost) error
	CreateMultiplePosts(multiCreatePost entity.MultiCreatePost) error
	GetAll() ([]entity.Post, error)
	GetByID(id uint64) (entity.Post, error)
	Update(post entity.UpdatePost) error
	Delete(id uint64) error
}

type postUsecase struct {
	repo  repository.PostRepository
	cache *infra.RedisClient
}

func NewPostUsecase(repo repository.PostRepository, cache *infra.RedisClient) PostUsecase {
	return &postUsecase{repo: repo, cache: cache}
}

func (p *postUsecase) Create(post entity.CreatePost) error {
	err := p.repo.Create(post)
	if err != nil {
		return err
	}

	p.cache.Delete("posts")
	return nil
}

func (p *postUsecase) GetAll() ([]entity.Post, error) {
	cachedPosts, err := p.cache.Get("posts")
	if err == nil && cachedPosts != "" {
		var posts []entity.Post
		json.Unmarshal([]byte(cachedPosts), &posts)
		return posts, nil
	}

	posts, err := p.repo.GetAll()
	if err != nil {
		return nil, err
	}

	cachedData, _ := json.Marshal(posts)
	p.cache.Set("posts", string(cachedData))
	return posts, nil
}

func (p *postUsecase) GetByID(id uint64) (entity.Post, error) {
	cachedPost, err := p.cache.Get("post:" + strconv.Itoa(int(id)))
	if err == nil && cachedPost != "" {
		var post entity.Post
		json.Unmarshal([]byte(cachedPost), &post)
		return post, nil
	}

	post, err := p.repo.GetByID(id)
	if err != nil {
		return post, err
	}

	cachedData, _ := json.Marshal(post)
	p.cache.Set("post:"+strconv.Itoa(int(id)), string(cachedData))
	return post, nil
}

func (p *postUsecase) Update(post entity.UpdatePost) error {
	existingPost, err := p.repo.GetByID(post.ID)
	if err != nil {
		return err
	}

	if existingPost.ID == 0 {
		return errors.New("not found")
	}

	err = p.repo.Update(post)
	if err != nil {
		return err
	}

	p.cache.Delete("post:" + strconv.Itoa(int(post.ID)))
	p.cache.Delete("posts")

	return nil
}

func (p *postUsecase) Delete(id uint64) error {
	err := p.repo.Delete(id)
	if err != nil {
		return err
	}

	p.cache.Delete("post:" + strconv.Itoa(int(id)))
	p.cache.Delete("posts")
	return nil
}

func (p *postUsecase) CreateMultiplePosts(multiCreatePost entity.MultiCreatePost) error {
	var posts []entity.Post

	for _, p := range multiCreatePost.Posts {
		posts = append(posts, entity.Post{
			Title:   p.Title,
			Content: p.Content,
			UserID:  multiCreatePost.UserID,
		})
	}

	return p.repo.CreatePosts(posts)
}
