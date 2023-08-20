package httphandler

import (
	"net/http"
	"strings"

	"github.com/aom31/GO-Inventory/config"
	"github.com/aom31/GO-Inventory/src/models"
	"github.com/aom31/GO-Inventory/src/repository"
	"github.com/labstack/echo/v4"
)

type UserHttpHandler struct {
	Cfg            *config.Config
	UserRepository *repository.UserRepository
}

func (handler *UserHttpHandler) FindOneUser(ctx echo.Context) error {
	ctxRequest := ctx.Request().Context()

	userId := strings.Trim(ctx.Param("userId"), " ")

	user, err := handler.UserRepository.FindOneUser(ctxRequest, userId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.Error{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, user)
}
