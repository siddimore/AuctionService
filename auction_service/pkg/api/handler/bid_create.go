package handlers

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/Fief-Finance/auction/models"
)

// Bid Create Handler
func (h* Handler) BidCreate(c echo.Context) error {

	auction := new(model.Bid)
  if err := c.Bind(auction); err != nil {
    return c.JSON(http.StatusBadRequest, "bad request")
  }

	// TODO Add Validation Code
	// if err = c.Validate(auction); err != nil {
	// 	return err
	// }
	ctx := c.Request().Context()
	err := h.BidService.Create(ctx, auction)

	if err != nil {
		c.JSON(apperrors.Status(err), err)
	}

	return c.JSON(http.StatusOK, nil)
}