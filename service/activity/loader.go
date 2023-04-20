package activity

import (
	"sync"

	"github.com/fatram/devcode-golang/domain/repository/mysql"
	"github.com/fatram/devcode-golang/pkg/genlog"
)

var (
	activityService     *ActivityService
	onceActivityService sync.Once
)

func LoadActivityService(logger genlog.Logger) *ActivityService {
	onceActivityService.Do(func() {
		activityService = &ActivityService{
			logger:     logger,
			repository: mysql.LoadActivityRepository(logger),
		}
	})
	return activityService
}
