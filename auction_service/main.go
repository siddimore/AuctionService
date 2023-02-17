package main

import (
	"github.com/Fief-Finance/auction/common"
	"github.com/labstack/echo"
	"github.com/Fief-Finance/auction/repository"
	"github.com/Fief-Finance/auction/service"
	"log"
)

type Response struct {
}

func main() {
	// Echo instance
	e := echo.New()

	ds, err := dbaccessor.InitDB()

	if err != nil {
		log.Fatalf("Unable to initialize data sources: %v\n", err)
	}

	// Create Repository Instances
	auctionRepository := repository.NewAuctionRepository(ds.DB)
	bidRepository := repository.NewBidRepository(ds.DB)

	// Create Services
	auctionService := service.NewAuctionService(&service.AuctionConfig{
		AuctionRepository: auctionRepository,
	})

	bidService := service.NewBidService(&service.BidConfig{
		BidRepository: bidRepository,
	})

	err = Inject(e, auctionService, bidService)

	if err != nil {
		log.Fatalf("Unable to initialize data sources: %v\n", err)
	}

}
