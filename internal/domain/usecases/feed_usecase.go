package usecases

import (
	"time"

	"github.com/prcryx/raft-server/internal/domain/entities"
)

type IFeedUseCase interface {
	GetFeeds() ([]entities.Feed, error)
}

type FeedUseCase struct {
	// repo repositories.FeedRepository
}

var _ IFeedUseCase = (*FeedUseCase)(nil)

func NewFeedUseCase() *FeedUseCase {
	return &FeedUseCase{
		// repo: repo,
	}
}

func (usecase *FeedUseCase) GetFeeds() ([]entities.Feed, error) {

	return []entities.Feed{
		{
			ID:        0,
			Body:      "Hi 0",
			CreatedAt: time.Now(),
		},
		{
			ID:        1,
			Body:      "Hi 1",
			CreatedAt: time.Now(),
		},
		{
			ID:        2,
			Body:      "Hi 2",
			CreatedAt: time.Now(),
		},
	}, nil
}
