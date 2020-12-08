package spot

import (
	"context"

	"flag"

	"github.com/walterjwhite/go/lib/external/spot/client"

	"path/filepath"
	"time"
)

const (
	// 2.5 minutes (if we go less than that, spot will block us and we don't want that)
	minRefreshInterval = time.Duration(150 * time.Second)
)

var (
	dataPath     = flag.String("spot-log-path", "~/.data/spot", "Directory to log spot tracking data")
	testDataPath = flag.String("test-spot-data-path", "", "Directory to use for test spot tracking data")
)

func (c *Configuration) Monitor(ctx context.Context) {
	c.ctx = ctx
	c.initActions()
	c.initFetcher()

	c.fetch()
}

func (c *Configuration) initFetcher() {
	if len(*testDataPath) == 0 {
		c.feedFetcher = &client.Feed{Id: c.Session.FeedId}
		return
	}

	c.feedFetcher = &client.FileFeed{Filename: filepath.Join(*testDataPath, c.Session.FeedId)}
}
