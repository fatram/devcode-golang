package todo

import (
	"sync"

	"github.com/fatram/devcode-golang/domain/repository/mysql"
	"github.com/fatram/devcode-golang/pkg/genlog"
)

var (
	todoService     *TodoService
	onceTodoService sync.Once
)

func LoadTodoService(logger genlog.Logger) *TodoService {
	onceTodoService.Do(func() {
		todoService = &TodoService{
			logger:     logger,
			repository: mysql.LoadTodoRepository(logger),
		}
	})
	return todoService
}
