# Unit Converter (Web)

This project was built as part of the  
[Unit Converter project](https://roadmap.sh/projects/unit-converter) on roadmap.sh.

The goal was to practice building a simple web application in Go using the standard library, focusing on HTTP handling,
server-side rendering, and basic production practices.

The application allows users to convert between different units of measurement via HTML forms and server-side
processing.

## Features

- Unit conversion for:
    - Length
    - Weight
    - Temperature
- Server-side rendered HTML (`html/template`)
- Form-based input (GET / POST)
- No database
- Graceful shutdown (SIGINT / SIGTERM)
- Basic request logging and panic recovery middleware

## Tech Stack

- Go
- `net/http`
- `html/template`
- Standard library only

## How to run

```bash
go run ./cmd/unitconv
