package handlers

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/google/uuid"
	"github.com/Fief-Finance/auction/errors"
)

// Auction Get Handler
func (h* Handler) AuctionGet(c echo.Context) error {

  uidString := c.Param("uid")
	ctx := c.Request().Context()
	uid, err := uuid.FromBytes([]byte(uidString))
  if err != nil {
    return c.JSON(http.StatusBadRequest, "bad request")
  }

	auction, err := h.AuctionService.Get(ctx, uid)

	if err != nil {
		c.JSON(apperrors.Status(err), err)
	}

	return c.JSON(http.StatusOK, auction)
}