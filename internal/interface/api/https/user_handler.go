package https

//Controller

import (
	"net/http"

	"simple-project/internal/application"
	"simple-project/internal/domain"

	"github.com/labstack/echo"
)

type UserHandler struct {
	UserUsecase application.UserUsecase
}

func NewUserHandler(uc application.UserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: uc,
	}
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.UserUsecase.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": "error read data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success read data",
		"results": users,
	})
}

func (h *UserHandler) AddUser(c echo.Context) error {
	var newUser domain.User
	err := c.Bind(&newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "error bind data: " + err.Error(),
		})
	}

	err = h.UserUsecase.AddUser(newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": "error insert data " + err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "success add user",
	})
}
