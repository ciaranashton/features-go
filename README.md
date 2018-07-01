# features-go
A simple features API for adding and removing features on a global level, as well as, an application by application basis.

This mirco-service is to explore building secure, robust and well tested APIs using idiomatic Golang.

## Setup
```bash
$ go run main.go
```

## Running tests
```bash
$ go test ./features -v
```

## TODOs
- PUT endpoint for updating a featue
- Allow features to be updated for a single application
- Unit testing
- Validation
- Authentication
- Trace Id middleware