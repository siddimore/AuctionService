package model

import (
	"github.com/google/uuid"
	"time"
)

type Auction struct {
	UID             uuid.UUID  `json:"uid"`
	Item            string `json:"Item"`
	MinBid          string `json:"Minbid"`
	MaxBid          string `json:"Maxbid"`
	AuctionExpiry   time.Duration `json:"AuctionExpiry"`
	BidIncrement    string `json:"BidIncrement"`
}