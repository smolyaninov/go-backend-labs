package domain

import (
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	ID        string          `json:"id"`
	Type      string          `json:"type"`
	Repo      Repo            `json:"repo"`
	CreatedAt time.Time       `json:"created_at"`
	Payload   json.RawMessage `json:"payload"`
}

type Repo struct {
	Name string `json:"name"`
}

func (e Event) Describe() string {
	switch e.Type {
	case "CreateEvent":
		return describeCreate(e)

	case "PushEvent":
		return describePush(e)

	case "IssuesEvent":
		return fmt.Sprintf("Opened a new issue in %s", e.Repo.Name)

	case "IssueCommentEvent":
		return fmt.Sprintf("Commented on an issue in %s", e.Repo.Name)

	case "PullRequestEvent":
		return fmt.Sprintf("Opened a pull request in %s", e.Repo.Name)

	case "WatchEvent":
		return fmt.Sprintf("Starred %s", e.Repo.Name)

	case "ForkEvent":
		return fmt.Sprintf("Forked %s", e.Repo.Name)

	default:
		return fmt.Sprintf("%s in %s", e.Type, e.Repo.Name)
	}
}

func describeCreate(e Event) string {
	var payload struct {
		Ref     string `json:"ref"`
		RefType string `json:"ref_type"`
	}
	if err := json.Unmarshal(e.Payload, &payload); err != nil {
		return fmt.Sprintf("Created something in %s", e.Repo.Name)
	}

	switch payload.RefType {
	case "repository":
		return fmt.Sprintf("Created repository %s", e.Repo.Name)
	case "branch":
		return fmt.Sprintf("Created branch %s in %s", payload.Ref, e.Repo.Name)
	case "tag":
		return fmt.Sprintf("Created tag %s in %s", payload.Ref, e.Repo.Name)
	default:
		return fmt.Sprintf("Created %s in %s", payload.RefType, e.Repo.Name)
	}
}

func describePush(e Event) string {
	var payload struct {
		Size int `json:"size"`
	}
	if err := json.Unmarshal(e.Payload, &payload); err != nil {
		return fmt.Sprintf("Pushed commits to %s", e.Repo.Name)
	}
	return fmt.Sprintf("Pushed %d commits to %s", payload.Size, e.Repo.Name)
}
