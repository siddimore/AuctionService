package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Fief-Finance/auction/models"
	"github.com/Fief-Finance/auction/errors"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type pGAuctionRepository struct {
	DB *sql.DB
}

// NewAuctionRepository is a factory for initializing Auction Repositories
func NewAuctionRepository(db *sql.DB) model.AuctionRepository {
	return &pGAuctionRepository{
		DB: db,
	}
}

// Create reaches out to database SQL api
func (r *pGAuctionRepository) Create(ctx context.Context, u *model.Auction) error {
	
	// TODO: Fix Me
	query := ""

	if _, err := r.DB.ExecContext(ctx, query); err != nil {
		// check unique constraint
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
			fmt.Printf("Could not create a Auction Reason: %v\n", err.Code.Name())
			return apperrors.NewConflict("AuctionId", u.UID.String())
		}
		return apperrors.NewInternal()
	}
	return nil
}

// FindByID fetches by id
func (r *pGAuctionRepository) FindByID(ctx context.Context, uid uuid.UUID) (*model.Auction, error) {
	auction := &model.Auction{}
	// TODO: Fix Me
	query := "SELECT * FROM auction WHERE uid=$1"

	if _, err := r.DB.ExecContext(ctx, query, uid); err != nil {
		return nil, apperrors.NewNotFound("uid", uid.String())
	}

	return auction, nil
}