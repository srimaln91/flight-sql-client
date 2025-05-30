# FlightSQL CLI

A simple CLI client written in Go to query Apache Arrow Flight SQL endpoints.

## Features

- Connects to a Flight SQL server (e.g., DuckDB, Dremio, custom backend)
- Executes SQL queries
- Prints results with headers in tabular format
- Supports both Linux and Windows builds

## Getting Started

### Requirements

- Go 1.20 or later
- Flight SQL-compatible server (e.g., DuckDB FlightSQL, Dremio)

### Build from Source

Clone the repository and build:

make build        # Builds for your current OS  
make build-linux  # Cross-compiles for Linux  
make build-win    # Cross-compiles for Windows

Output binaries will be placed in the project root or build/ directory (if using GitHub Actions).

## Usage

./flightsql-client --host=localhost --port=32010 --query="SELECT * FROM my_table"

### Available Flags

Flag        Description                                      Default  
--host      Flight SQL server hostname                       localhost  
--port      Flight SQL server port                           32010  
--query     SQL query to execute                             (required)  
--tls       Enable TLS (not implemented yet)                 false  
--timeout   Query timeout duration (e.g. 10s, 5m)            10s

### Example

./flightsql-client --host=127.0.0.1 --port=32010 --query="SELECT id, name FROM users"

Output:  
id    name  
--    ------  
1     Alice  
2     Bob

## GitHub Actions

On push or PR to main, binaries for:

- linux/amd64
- windows/amd64

are built and uploaded as downloadable artifacts.

## TODO

- TLS support  
- Authentication support  
- Output as CSV/JSON  
- Query from file  
- Docker support

## Contributing

Feel free to open issues or PRs. To contribute:

go fmt ./...  
go build -v .

## License

MIT License.