package httphandler

import (
	"net/http"
	"strings"

	"github.com/aom31/GO-Inventory/config"
	"github.com/aom31/GO-Inventory/src/models"
	"github.com/aom31/GO-Inventory/src/service"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHttpHandler struct {
	Cfg         *config.Config
	Client      *mongo.Client
	UserService service.IUserService
}

// UserHttpHandler is initial values of user http handler.
func NewUserHttpHandler(Cfg *config.Config, Client *mongo.Client) *UserHttpHandler {
	return &UserHttpHandler{
		Cfg:         Cfg,
		UserService: service.NewUserService(Client),
	}
}

func (handler *UserHttpHandler) FindOneUser(ctx echo.Context) error {
	ctxRequest := ctx.Request().Context()

	userId := strings.Trim(ctx.Param("userId"), " ")

	user, err := handler.UserService.FindUserById(ctxRequest, userId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.Error{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, user)
}
