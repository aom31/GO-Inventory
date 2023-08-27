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

type ItemHttpHandler struct {
	Cfg         *config.Config
	Client      *mongo.Client
	ItemService service.IItemService
}

// UserHttpHandler is initial values of user http handler.
func NewItemHttpHandler(Cfg *config.Config, Client *mongo.Client) *ItemHttpHandler {
	return &ItemHttpHandler{
		Cfg:         Cfg,
		ItemService: service.NewItemService(Client),
	}
}

func (handle *ItemHttpHandler) FindItems(c echo.Context) error {
	ctx := c.Request().Context()
	items := make([]models.Item, 0)

	if err := handle.ItemService.FindItems(ctx, &items); err != nil {
		return c.JSON(http.StatusBadRequest, &models.Error{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, items)
}

func (handle *ItemHttpHandler) FindOneItem(c echo.Context) error {
	ctx := c.Request().Context()

	itemId := strings.Trim(c.Param("itemId"), " ")
	item, err := handle.ItemService.FindOneItem(ctx, itemId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.Error{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, item)
}
