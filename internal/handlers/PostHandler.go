package handlers

import (
	"User-Post-Backend/infra"
	"User-Post-Backend/internal/entity"
	"User-Post-Backend/internal/helpers"
	"User-Post-Backend/internal/repository"
	"User-Post-Backend/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PostHandler struct {
	postUsecase usecase.PostUsecase
}

func NewPostHandler(app *fiber.App, db *gorm.DB, cache *infra.RedisClient) {
	repo := repository.NewPostRepository(db)
	usecase := usecase.NewPostUsecase(repo, cache)
	handler := &PostHandler{postUsecase: usecase}

	apiv1 := app.Group("/api/v1")

	apiv1.Post("/posts", handler.Create)
	app.Post("/multi-posts", handler.CreateMultiplePosts)
	apiv1.Get("/posts", handler.GetAll)
	apiv1.Get("/posts/:id", handler.GetByID)
	apiv1.Put("/posts/:id", handler.Update)
	apiv1.Delete("/posts/:id", handler.Delete)
}

// @Summary Get all posts
// @Description Get a list of all posts
// @Tags posts
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Post
// @Router /api/v1/posts [get]
func (h *PostHandler) GetAll(c *fiber.Ctx) error {
	posts, err := h.postUsecase.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SendResponse(c, fiber.StatusOK, "successfully retrieved posts", posts)
}

// @Summary Get a post by ID
// @Description Get a single post by ID
// @Tags posts
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Success 200 {object} entity.Post
// @Router /api/v1/posts/{id} [get]
func (h *PostHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid ID")
	}
	post, err := h.postUsecase.GetByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SendResponse(c, fiber.StatusOK, "successfully retrieved post", post)
}

// @Summary Create a new post
// @Description Add a new post to the database
// @Tags posts
// @Accept  json
// @Produce  json
// @Param post body entity.CreatePost true "Post data"
// @Success 201 {string} string "Post created"
// @Router /api/v1/posts [post]
func (h *PostHandler) Create(c *fiber.Ctx) error {
	var post entity.CreatePost
	if err := c.BodyParser(&post); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := h.postUsecase.Create(post); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SendResponse(c, fiber.StatusCreated, "post created successfully", nil)
}

// @Summary Create a new multi post
// @Description Add a new post to the database
// @Tags posts
// @Accept  json
// @Produce  json
// @Param post body entity.MultiCreatePost true "Post data"
// @Success 201 {string} string "Post created"
// @Router /api/v1/multi-posts [post]
func (h *PostHandler) CreateMultiplePosts(c *fiber.Ctx) error {
	var multiCreatePost entity.MultiCreatePost
	if err := c.BodyParser(&multiCreatePost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.postUsecase.CreateMultiplePosts(multiCreatePost); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create posts"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Posts created successfully"})
}

// @Summary Update an existing post
// @Description Update a post's data
// @Tags posts
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Param post body entity.Post true "Post data"
// @Success 200 {string} string "Post updated"
// @Router /api/v1/posts/{id} [put]
func (h *PostHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid ID")
	}

	var post entity.UpdatePost
	if err := c.BodyParser(&post); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	post.ID = id
	if err := h.postUsecase.Update(post); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SendResponse(c, fiber.StatusOK, "post updated successfully", nil)
}

// @Summary Update an existing post
// @Description Update a post's data
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Param post body entity.UpdatePost true "Post data"
// @Success 200 {object} helpers.StandardResponse "Post updated successfully"
// @Failure 400 {object} helpers.StandardResponse "Invalid ID or request body"
// @Failure 404 {object} helpers.StandardResponse "Data not found"
// @Router /api/v1/posts/{id} [patch]
func (h *PostHandler) UpdatePatch(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid ID")
	}

	var req entity.UpdatePost
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	req.ID = id
	err = h.postUsecase.Update(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SendResponse(c, fiber.StatusOK, "Post updated successfully", nil)
}

// @Summary Delete a post
// @Description Delete a post by ID
// @Tags posts
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Success 200 {string} string "Post deleted"
// @Router /api/v1/posts/{id} [delete]
func (h *PostHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid ID")
	}
	if err := h.postUsecase.Delete(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SendResponse(c, fiber.StatusOK, "post deleted successfully", nil)
}
