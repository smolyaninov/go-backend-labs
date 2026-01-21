package service

import (
	"go-backend-labs/01-beginner/github-user-activity/internal/domain"
)

type ActivityService interface {
	FilterEventsByType(events []domain.Event, eventType string) []domain.Event
}

type activityService struct{}

func NewActivityService() ActivityService {
	return &activityService{}
}

func (s *activityService) FilterEventsByType(events []domain.Event, eventType string) []domain.Event {
	if eventType == "" {
		return events
	}

	var filteredEvents []domain.Event
	for _, event := range events {
		if event.Type == eventType {
			filteredEvents = append(filteredEvents, event)
		}
	}
	return filteredEvents
}
