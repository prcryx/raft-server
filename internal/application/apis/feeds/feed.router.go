package feed

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/prcryx/raft-server/internal/common/constants/routesconst"
)

func FeedRouter(router chi.Router, fc IFeedController) {
	router.Get(fmt.Sprintf("%v/{user-id}", routesconst.UserFeeds), fc.GetFeeds)
}
