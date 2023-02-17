package main

import (
	"github.com/Fief-Finance/auction/models"
	"github.com/Fief-Finance/auction/pkg/api/handler"
	"github.com/labstack/echo"
	"os"
	"time"
)

func Inject(e *echo.Echo, auctionService model.AuctionService, bidService model.BidService) error {
	router := e
	baseURL := os.Getenv("ACCOUNT_API_URL")
	handlerTimeout, _ := time.ParseDuration(os.Getenv("HANDLER_TIMEOUT"))

	handlers.SetUp(&handlers.Config{
		R:               router,
		AuctionService:  auctionService,
		BidService:      bidService,
		BaseURL:         baseURL,
		TimeoutDuration: handlerTimeout})

	return nil
}
