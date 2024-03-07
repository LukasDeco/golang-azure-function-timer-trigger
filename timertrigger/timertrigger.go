package embeddingsgen

import (
	"context"
	"time"

	"github.com/golang/glog"
)

func Run(ctx context.Context) {
	glog.Info("Timer trigger function executed at: %s", time.Now().Format(time.RFC3339))
}
