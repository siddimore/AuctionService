package handlers

import (
	"net/http"
	"github.com/labstack/echo"
)

// Bid Get Handler
func (h* Handler) BidGet(c echo.Context) error {
	uidString := c.Param("uid")
	ctx := c.Request().Context()
	uid, err := uuid.FromBytes([]byte(uidString))
  if err != nil {
    return c.JSON(http.StatusBadRequest, "bad request")
  }

	bid, err := h.BidService.Get(ctx, uid)

	if err != nil {
		c.JSON(apperrors.Status(err), err)
	}
	return c.JSON(http.StatusOK, bid)
}