package model

import (
	"context"
	"github.com/google/uuid"
)

type AuctionService interface {
	Get(ctx context.Context, uid uuid.UUID) (*Auction, error)
	Create(ctx context.Context, u *Auction) error
}

type BidService interface {
	Get(ctx context.Context, uid uuid.UUID) (*Bid, error)
	Create(ctx context.Context, u *Bid) error
}

// AuctionRepository defines methods the service layer expects
type AuctionRepository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*Auction, error)
	Create(ctx context.Context, u *Auction) error
	// Update(ctx context.Context, u *Auction) error
}

type BidRepository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*Bid, error)
	Create(ctx context.Context, u *Bid) error
	// Update(ctx context.Context, u *Bid) error
}