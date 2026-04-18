# Kids Typing Practice

A typing game for kids built with Go and HTMX. A sentence appears on screen — type it correctly and a funny video clip plays as a reward.

## Stack

- **Go** — HTTP server, sentence + clips API
- **HTMX** — frontend interactivity, no JS framework
- Static HTML/CSS frontend

## How to run

**Requirements:** Go 1.18+

```bash
git clone https://github.com/nerryc/KidsTypingPractice.git
cd KidsTypingPractice
go run main.go
```

Open [http://localhost:8080](http://localhost:8080).

## Adding sentences

Open `main.go` and add strings to the `sentences` slice:

```go
var sentences = []string{
    "The cat sat on the mat.",
    "Your new sentence here.",
    // ...
}
```

## Adding video clips

Drop `.mp4` files into `static/clips/`. They are served automatically — no code changes needed. A random clip plays each time a sentence is typed correctly.

## Project structure

```
KidsTypingPractice/
├── main.go          # HTTP server + /sentence and /clips API
├── go.mod
└── static/
    ├── index.html   # Game UI
    └── clips/       # Add .mp4 reward videos here
```

## API

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/sentence?index=N` | GET | Returns sentence at index N + next index |
| `/clips` | GET | Returns list of available `.mp4` clip filenames |
