package util

import (
	"context"

	log "github.com/sirupsen/logrus"
)

func Action(ctx context.Context, action string) *log.Entry {
	return log.WithContext(ctx).WithField("action", action)
}
