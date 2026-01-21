package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"go-backend-labs/01-beginner/github-user-activity/internal/client"
	"go-backend-labs/01-beginner/github-user-activity/internal/domain"
	"go-backend-labs/01-beginner/github-user-activity/internal/repository"
	"go-backend-labs/01-beginner/github-user-activity/internal/service"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: github-activity <username> [--type=EventType] [--no-cache]")
		os.Exit(1)
	}

	username := os.Args[1]
	var eventType string
	noCache := false

	for _, arg := range os.Args[2:] {
		switch {
		case strings.HasPrefix(arg, "--type="):
			eventType = strings.TrimPrefix(arg, "--type=")
		case arg == "--no-cache":
			noCache = true
		}
	}

	apiClient := client.NewGitHubClient()
	cache := repository.NewJSONCache[domain.Event](".cache", 10*time.Minute)
	svc := service.NewActivityService()

	var events []domain.Event
	var err error

	if !noCache {
		events, err = cache.Load(username)
	}

	if noCache || err != nil {
		events, err = apiClient.GetUserEvents(username)
		if err != nil {
			log.Fatalf("Error fetching events: %v", err)
		}
		if !noCache {
			_ = cache.Save(username, events)
		}
	}

	filteredEvents := svc.FilterEventsByType(events, eventType)

	for _, event := range filteredEvents {
		fmt.Printf(
			"[%s] %s\n",
			event.CreatedAt.Format("2006-01-02 15:04"),
			event.Describe(),
		)
	}
}
