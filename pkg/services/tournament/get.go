package tournament

import (
	"context"
	"fmt"
	"strings"

	"github.com/jairogloz/go-l/pkg/domain"
)

// Get retrieves a tournament by ID from repository layer.
func (s *Service) Get(ctx context.Context, id string) (tournament *domain.Tournament, err error) {

	id = strings.TrimSpace(id)
	if id == "" {
		return nil, fmt.Errorf("error getting a tournament %w,", domain.ErrInvalidParam)
	}

	tournament, err = s.Repo.Get(ctx, id)
	if err != nil {
		return nil, domain.ManageError(err, "error getting a tournament")
	}

	return tournament, nil
}
