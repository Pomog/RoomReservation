#!/bin/bash
go build -o bookings cmd/web/*.go && 
./bookings -dbname=bookings -dbuser=thoryur -dbpassword=plot123123

#chmod +x run.sh
# go test -v ./...
# go test -v -cover ./...
# go test -v -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
# go test -v -coverprofile=coverage.out ./... && go tool cover -func=coverage.out
# go test -v -coverprofile=coverage.out ./... && go tool cover -func=coverage.out | grep -i total
# ./run.sh

# sudo service postgresql status
# sudo -u postgres psql
# install soda brew install gobuffalo/tap/pop
# 802
# ~/go/bin/MailHog