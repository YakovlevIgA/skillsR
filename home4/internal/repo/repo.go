package repo

import (
	"context"
	"go.uber.org/zap"
	"sync"
	"github.com/pkg/errors"
)

// Слой репозитория, здесь должны быть все методы, связанные с базой данных

// SQL-запрос на вставку задачи
//const (
//	insertTaskQuery = `INSERT INTO tasks (title, description) VALUES ($1, $2) RETURNING id;`
//	getTaskQuery    = `SELECT * FROM tasks WHERE id = $1;`
//)

type InMemoryRepository struct {
	mu     sync.RWMutex
	tasks  map[int]Task
	nextID int
	logger *zap.Logger // Добавляем логгер
}

// NewInMemoryRepository теперь принимает логгер
func NewInMemoryRepository(logger *zap.Logger) *InMemoryRepository {
	return &InMemoryRepository{
		tasks:  make(map[int]Task),
		nextID: 1,
		logger: logger.With(zap.String("component", "inMemoryRepository")), // Добавляем контекст
	}
}

// Repository - интерфейс с методом создания задачи
type Repository interface {
	CreateTask(ctx context.Context, task Task) (int, error)
	GetTask(ctx context.Context, id int) (Task, error)
	GetTasks(ctx context.Context) ([]Task, error)
	UpdateTask(ctx context.Context, id int, task Task) error
	DeleteTask(ctx context.Context, id int) error
}

func (r *InMemoryRepository) CreateTask(ctx context.Context, task Task) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	task.ID = r.nextID
	r.tasks[r.nextID] = task

	r.logger.Debug("Task created",
		zap.Int("id", task.ID),
		zap.String("title", task.Title),
	)

	r.nextID++
	return task.ID, nil
}

func (r *InMemoryRepository) GetTask(ctx context.Context, id int) (Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		r.logger.Warn("Task not found", zap.Int("id", id))
		return Task{}, errors.New("task not found")
	}

	r.logger.Debug("Task retrieved", zap.Int("id", id))
	return task, nil
}

func (r *InMemoryRepository) GetTasks(ctx context.Context) ([]Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	count := len(r.tasks)
	r.logger.Debug("Getting all tasks", zap.Int("count", count))

	result := make([]Task, 0, count)
	for _, task := range r.tasks {
		result = append(result, task)
	}
	return result, nil
}

func (r *InMemoryRepository) UpdateTask(ctx context.Context, id int, task Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; !exists {
		r.logger.Warn("Task not found for update", zap.Int("id", id))
		return errors.New("task not found")
	}

	task.ID = id
	r.tasks[id] = task

	r.logger.Info("Task updated",
		zap.Int("id", id),
		zap.String("new_status", task.Status),
	)
	return nil
}

func (r *InMemoryRepository) DeleteTask(ctx context.Context, id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; !exists {
		r.logger.Warn("Task not found for deletion", zap.Int("id", id))
		return errors.New("task not found")
	}

	delete(r.tasks, id)
	r.logger.Info("Task deleted", zap.Int("id", id))
	return nil
}
