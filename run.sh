#!/bin/bash
go build -o bookings cmd/web/*.go && ./bookings
```
#chmod +x run.sh
# go test -v ./...
# go test -v -cover ./...
# go test -v -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
# go test -v -coverprofile=coverage.out ./... && go tool cover -func=coverage.out
# go test -v -coverprofile=coverage.out ./... && go tool cover -func=coverage.out | grep -i total
# ./run.sh