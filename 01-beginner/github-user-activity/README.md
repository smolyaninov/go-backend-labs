# GitHub User Activity

This is my second project in Go, completed as part of the [GitHub User Activity project](https://roadmap.sh/projects/github-user-activity) on roadmap.sh.

The goal was to explore working with APIs, JSON parsing, and CLI tools in Go by building a practical command-line application.

The application is a CLI tool that takes a GitHub username and displays the user's recent public activity using the GitHub API (`/users/<username>/events`).

## Features

* Fetches event data from the GitHub API in JSON format
* Displays events in a human-readable format
* Filters events by type (`--type=PushEvent`)
* Caches results locally with a time-to-live (TTL)
* `--no-cache` option to force fresh API request

## Usage

```bash
github-activity <username>
github-activity <username> --type=PushEvent
github-activity <username> --no-cache
github-activity <username> --type=PushEvent --no-cache
```

## Build

```bash
go build -o github-activity ./cmd/github-activity
```
