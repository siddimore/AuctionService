package model

import (
	"github.com/google/uuid"
	"time"
)

type Bid struct {
	UID       uuid.UUID  `json:"uid"`
	AuctionId string `json:"AuctionId"`
	Item      string `json:"Item"`
	Bid       string `json:"Bid"`
	CreateTime time.Time `json:"CreateTime"`
	UpdateTime time.Time `json:"UpdateTime"`
	// TODO: Add UserDetails
}