# engram

> A lightweight, fast, and extensible knowledge memory engine for AI agents and developer tooling.

Fork of [Gentleman-Programming/engram](https://github.com/Gentleman-Programming/engram).

## Overview

Engram is a local-first memory and context management system designed to help AI assistants and developer tools maintain persistent, searchable knowledge across sessions. It stores information in compressed chunked JSONL files with a manifest index for fast retrieval.

## Features

- 🧠 **Persistent memory** — Store and retrieve context across sessions
- ⚡ **Fast chunk-based storage** — Compressed JSONL chunks for efficient I/O
- 🔍 **Semantic search** — Query stored knowledge by relevance
- 🔌 **Plugin marketplace** — Extend functionality via `.claude-plugin/marketplace.json`
- 📦 **Zero external dependencies** — Pure Go implementation

## Installation

```bash
go install github.com/your-org/engram@latest
```

Or build from source:

```bash
git clone https://github.com/your-org/engram.git
cd engram
go build ./...
```

## Usage

### CLI

```bash
# Store a memory chunk
engram store --content "Go interfaces are satisfied implicitly"

# Query memory
engram query --q "how do Go interfaces work"

# List all chunks
engram list

# Show manifest
engram manifest

# Delete a chunk by hash
engram delete --hash <hash>
```

### As a library

```go
import "github.com/your-org/engram/pkg/memory"

store, err := memory.NewStore(".engram")
if err != nil {
    log.Fatal(err)
}

// Write a chunk
err = store.Write(memory.Chunk{
    Content: "Go channels enable safe communication between goroutines",
    Tags:    []string{"go", "concurrency"},
})

// Query chunks
results, err := store.Query("goroutine communication")
```

## Storage Format

Engram stores data in `.engram/` directory:

```
.engram/
  manifest.json        # Index of all chunks with metadata
  chunks/
    <hash>.jsonl.gz    # Compressed JSONL chunk files
```

Each chunk file contains one JSON object per line, compressed with gzip.

> **Personal note:** I've been using `~/.engram` as the store path instead of a project-local `.engram/` so that memories persist globally across different project directories. Just pass `--store $HOME/.engram` to the CLI or set `ENGRAM_STORE=$HOME/.engram` in your shell profile.

## Plugin System

Plugins are configured in `.claude-plugin/marketplace.json`. See the marketplace file for available plugins and configuration options.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feat/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feat/amazing-feature`)
5. Open a Pull Request

Please use the [bug report template](.github/ISSUE_TEMPLATE/bug_report.yml) for issues.

## License

MIT License — see [LICENSE](LICENSE) for details.
