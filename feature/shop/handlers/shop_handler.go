package handlers

import "github.com/labstack/echo/v4"

type ShopHandler interface {
	DetectShop(context echo.Context) error
}
