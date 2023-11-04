package feed

import (
	"net/http"

	"github.com/prcryx/raft-server/internal/common/utils"
	"github.com/prcryx/raft-server/internal/domain/usecases"
)

type IFeedController interface {
	GetFeeds(http.ResponseWriter, *http.Request)
}

type FeedController struct {
	feedUseCase *usecases.FeedUseCase
}

func NewFeedController(feedUseCase *usecases.FeedUseCase) *FeedController {
	return &FeedController{
		feedUseCase: feedUseCase,
	}
}

var _ IFeedController = (*FeedController)(nil)

func (fc FeedController) GetFeeds(w http.ResponseWriter, r *http.Request) {
	feed, err := fc.feedUseCase.GetFeeds()
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseWithJSONData(w, http.StatusOK, feed)
}
