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

type pGBidRepository struct {
	DB *sql.DB
}

// NewBidRepository is a factory for initializing Bid Repositories
func NewBidRepository(db *sql.DB) model.BidRepository {
	return &pGBidRepository{
		DB: db,
	}
}

// Create reaches out to database SQL api
func (r *pGBidRepository) Create(ctx context.Context, u *model.Bid) error {
	
	// TODO: Fix Me
	query := ""

	if _, err := r.DB.ExecContext(ctx, query); err != nil {
		// check unique constraint
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
			fmt.Printf("Could not create a Bid Reason: %v\n", err.Code.Name())
			return apperrors.NewConflict("BidId", u.UID.String())
		}
		return apperrors.NewInternal()
	}
	return nil
}

// FindByID fetches by id
func (r *pGBidRepository) FindByID(ctx context.Context, uid uuid.UUID) (*model.Bid, error) {
	auction := &model.Bid{}
	// TODO: Fix Me
	query := "SELECT * FROM bid WHERE uid=$1"

	if _, err := r.DB.ExecContext(ctx, query, uid); err != nil {
		return nil, apperrors.NewNotFound("uid", uid.String())
	}

	return auction, nil
}