package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/Fief-Finance/auction/models"
)

// auctionService acts as a struct for injecting an implementation of AuctionRepository
// for use in service methods
type auctionService struct {
	AuctionRepository  model.AuctionRepository
}

type AuctionConfig struct {
	AuctionRepository  model.AuctionRepository
}

func NewAuctionService(c *AuctionConfig) model.AuctionService {
	return &auctionService{ 
		AuctionRepository:  c.AuctionRepository,
	}
}

// Get retrieves a user based on their uuid
func (s *auctionService) Get(ctx context.Context, uid uuid.UUID) (*model.Auction, error) {
	u, err := s.AuctionRepository.FindByID(ctx, uid)

	return u, err
}

func (s *auctionService) Create(ctx context.Context, u *model.Auction) error {

	err := s.AuctionRepository.Create(ctx, u)

	if (err != nil){
		return err
	}

	return nil
}