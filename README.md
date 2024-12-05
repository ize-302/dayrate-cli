# Dayrate CLI
Simple CLI tool to rate your day. Built while learning Go.

## Development
Make sure you cd into `cmd/dayrate-cli`

##### Run in development
`go run . list` - list all ratings
`go run . add` - add today's rating
`go run . help` - help menu

## Build
From `cmd/dayrate-cli`
`go build -C . -o ../../bin/APP_NAME`

##### Run build
`APP_NAME list` - list all ratings
`APP_NAME add` - add today's rating
`APP_NAME help` - help menu
