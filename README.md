# Geminix

Geminix is a Go-based project for interacting with Gemini protocol servers.

## Project Structure

```
go.mod
go.sum
LICENSE
README.md
cmd/
    geminix/
        main.go
internal/
    prompt/
        handler.go
pkg/
    gemini/
        client.go
```

- **cmd/geminix/main.go**: Entry point for the CLI application.
- **internal/prompt/handler.go**: Handles prompt logic for user input.
- **pkg/gemini/client.go**: Contains the Gemini protocol client implementation.

## Installation

To install Geminix using `go install`, run:

```sh
go install github.com/alpernae/geminix/cmd/geminix@latest
```

## Building

To build the project manually, run:

```sh
go build -o geminix ./cmd/geminix
```

## Usage

After building or installing, run the CLI:

```sh
geminix
```

## License

This project is licensed under the GNU General Public License v3.0. See [LICENSE](LICENSE) for details.