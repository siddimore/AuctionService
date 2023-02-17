package handlers

import (
	"net/http"
	"github.com/labstack/echo"
)

// Auction Close Handler
func (h* Handler) AuctionClose(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}