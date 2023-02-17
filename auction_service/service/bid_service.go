package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/Fief-Finance/auction/models"
)

// auctionService acts as a struct for injecting an implementation of BidRepository
// for use in service methods
type bidService struct {
	BidRepository  model.BidRepository
}

type BidConfig struct {
	BidRepository  model.BidRepository
}

func NewBidService(c *BidConfig) model.BidService {
	return &bidService{
		BidRepository:  c.BidRepository,
	}
}

// Get retrieves a user based on their uuid
func (s *bidService) Get(ctx context.Context, uid uuid.UUID) (*model.Bid, error) {
	u, err := s.BidRepository.FindByID(ctx, uid)

	return u, err
}

func (s *bidService) Create(ctx context.Context, u *model.Bid) error {

	err := s.BidRepository.Create(ctx, u)

	if (err != nil){
		return err
	}

	return nil
}