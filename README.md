# home_assignment

### Run service with database and demo data:

- create a `.env` file and fill out with data (see `example.env`)
- run `docker-compose -f compose.yml up --build -d`

### Run business logic tests:

- run `go test ./...`