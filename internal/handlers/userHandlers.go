package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"task-api/internal/userService"
	"task-api/internal/web/users"
)

type UserHandlers struct {
	svc *userService.Service
}

func NewUserHandlers(svc *userService.Service) *UserHandlers {
	return &UserHandlers{svc: svc}
}

func (h *UserHandlers) GetUsers(c echo.Context) error {
	list, err := h.svc.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, list)
}

func (h *UserHandlers) PostUser(c echo.Context) error {
	var body users.UserCreate
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	user, err := h.svc.CreateUser(body.Email, body.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}

// Дописать ручки и новый мэйн и начать тестировать как  все будет работать
func (h *UserHandlers) PatchUserByID(c echo.Context, id uint64) error {
	var body users.UserUpdate
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	email := ""
	if body.Email != nil {
		email = *body.Email
	}
	password := ""
	if body.Password != nil {
		password = *body.Password
	}

	updated, err := h.svc.UpdateUser(uint(id), email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updated)
}

func (h *UserHandlers) DeleteUserByID(c echo.Context, id uint64) error {
	if err := h.svc.DeleteUser(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
