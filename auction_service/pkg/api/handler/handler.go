package handlers

import (
	"github.com/labstack/echo"
	"github.com/Fief-Finance/auction/models"
	"time"
)

// Handler struct holds required services for handler to function
type Handler struct{
  AuctionService model.AuctionService
	BidService     model.BidService
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
    R               *echo.Echo
		AuctionService model.AuctionService
		BidService     model.BidService
		BaseURL         string
		TimeoutDuration time.Duration
}


func SetUp(c *Config) {
	h := &Handler{
		AuctionService: c.AuctionService,
		BidService: c.BidService,
	}

	auctionApis := c.R.Group(c.BaseURL)
	// TODO: Add Auth Middleware
	auctionApis.GET("/auction/:uid", h.AuctionGet)
	auctionApis.PUT("/auction/create", h.AuctionCreate)
	auctionApis.PUT("/auction/close", h.AuctionClose)
	auctionApis.POST("/bid", h.BidCreate)
	auctionApis.GET("/bid/:uid", h.BidGet)

}