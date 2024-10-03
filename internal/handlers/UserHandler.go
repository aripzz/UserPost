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

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(app *fiber.App, db *gorm.DB, cache *infra.RedisClient) {
	repo := repository.NewUserRepository(db)
	usecase := usecase.NewUserUsecase(repo, cache)
	handler := &UserHandler{userUsecase: usecase}

	apiv1 := app.Group("/api/v1")

	apiv1.Post("/users", handler.Create)
	apiv1.Get("/users", handler.GetAll)
	apiv1.Get("/users/:id", handler.GetByID)
	apiv1.Put("/users/:id", handler.Update)
	apiv1.Delete("/users/:id", handler.Delete)
}

// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.User
// @Router /api/v1/users [get]
func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.userUsecase.GetAll()
	if err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SendResponse(c, fiber.StatusOK, "successfully", users)
}

// @Summary Get a user by ID
// @Description Get a single user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} entity.User
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid ID")
	}
	user, err := h.userUsecase.GetByID(id)
	if err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SendResponse(c, fiber.StatusOK, "successfully", user)
}

// @Summary Create a new user
// @Description Add a new user to the database
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body entity.CreateUser true "User data"
// @Success 201 {string} string "User created"
// @Router /api/v1/users [post]
func (h *UserHandler) Create(c *fiber.Ctx) error {
	var user entity.CreateUser
	if err := c.BodyParser(&user); err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if err := h.userUsecase.Create(user); err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SendResponse(c, fiber.StatusCreated, "successfully", nil)
}

// @Summary Update an existing user
// @Description Update a user's data
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body entity.User true "User data"
// @Success 200 {string} string "User updated"
// @Router /api/v1/users/{id} [put]
func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid ID")
	}

	var user entity.User
	if err := c.BodyParser(&user); err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	user.ID = id
	if err := h.userUsecase.Update(user); err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SendResponse(c, fiber.StatusOK, "update successfully", nil)
}

// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {string} string "User deleted"
// @Router /api/v1/users/{id} [delete]
func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid ID")
	}
	if err := h.userUsecase.Delete(id); err != nil {
		return helpers.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return helpers.SendResponse(c, fiber.StatusOK, "deleted successfully", nil)
}
