# test-redis
Practice, Go support for Redis.

Simple operations:
- open DB connection;
- write a value;
- read the value.

## Dependencies
- install Go
- run Redis server like the following: ` docker run -d -p 6379:6379 redislabs/redismod`
## Build & run

Update vendor:
`go mod tidy`

Then just run `make` to run or `make build` to make binary (placed to "bin" folder).
