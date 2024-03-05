package handlers

import (
	"awesomeProject/feature/shop/models"
	"awesomeProject/feature/shop/usecases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type shopHttpHandler struct {
	shopUseCase usecases.ShopUseCase
}

func NewShopHttpHandler(shopUseCase usecases.ShopUseCase) ShopHandler {
	return &shopHttpHandler{
		shopUseCase: shopUseCase,
	}
}

func (h *shopHttpHandler) DetectShop(c echo.Context) error {
	reqBody := new(models.AddShopData)

	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, http.StatusBadRequest, "Bad request")
	}

	if err := h.shopUseCase.ShopDataProcessing(reqBody); err != nil {
		return response(c, http.StatusInternalServerError, "Processing data failed")
	}

	return response(c, http.StatusOK, "Success 🪳🪳🪳")
}
