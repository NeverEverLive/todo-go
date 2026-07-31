package statistics_service

import (
	"context"
	"fmt"
	"time"

	"github.com/NeverEverLive/todo-go/internal/core/domain"
	core_errors "github.com/NeverEverLive/todo-go/internal/core/errors"
)

func (s *StatisticsService) GetStatistics(
	ctx context.Context,
	userID *string,
	dateFrom *time.Time,
	dateTo *time.Time,
) (domain.Statistics, error) {
	if dateFrom != nil && dateTo != nil {
		if dateFrom.After(*dateTo) || dateFrom.Equal(*dateTo) {
			return domain.Statistics{}, fmt.Errorf(
				"`date_from` must be after `date_to`: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	tasks, err := s.statisticsRepository.GetTasks(ctx, userID, dateFrom, dateTo)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf(
			"get tasks from repository `%w`",
			err,
		)
	}

	statistics := calcStatistics(tasks)

	return statistics, nil
}

func calcStatistics(tasks []domain.Task) domain.Statistics {
	totalTasks := len(tasks)

	if totalTasks == 0 {
		return domain.NewStatistics(
			0,
			0,
			nil,
			nil,
		)
	}

	completedTasks := 0
	var summaryCompletionTasksTime time.Duration

	for _, task := range tasks {
		if task.Completed {
			completedTasks++
		}

		completionTasksTime := task.CompletionDuration()
		if completionTasksTime != nil {
			summaryCompletionTasksTime += *completionTasksTime
		}
	}

	tasksCompletedRate := float64(completedTasks) / float64(totalTasks) * 100
	var averageCompletionTimeTasks *time.Duration
	if completedTasks > 0 && summaryCompletionTasksTime != 0 {
		avg := summaryCompletionTasksTime / time.Duration(completedTasks)
		averageCompletionTimeTasks = &avg
	}

	return domain.NewStatistics(
		totalTasks,
		completedTasks,
		&tasksCompletedRate,
		averageCompletionTimeTasks,
	)
}
