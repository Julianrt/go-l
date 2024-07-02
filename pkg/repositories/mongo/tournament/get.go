package tournament

import (
	"context"
	"errors"

	"github.com/jairogloz/go-l/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Get retrieves a tournament by ID from collection tournaments.
func (r *Repository) Get(ctx context.Context, id string) (*domain.Tournament, error) {
	tournamentId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, domain.ErrInvalidParam
	}

	tournament := new(domain.Tournament)
	err = r.Collection.FindOne(ctx, bson.M{"_id": tournamentId}).Decode(tournament)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrNotFound
		}
		var commandErr *mongo.CommandError
		if errors.As(err, &commandErr) && commandErr.HasErrorLabel("NetworkTimeout") {
			return nil, domain.ErrTimeout
		}
		return nil, err
	}

	tournament.ID = tournamentId.Hex()
	return tournament, nil
}
