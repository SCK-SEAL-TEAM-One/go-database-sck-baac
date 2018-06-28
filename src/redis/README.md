# Go Database Connection

## How to Install Project
### Run with docker image
```bash
docker pull redis
docker run --name redis -d -p 6379:6379 redis
```

### View Redis Value Using Command Line
```bash
docker exec -it redis redis-cli
```
[See more redis command](https://redis.io/commands)
### Install go mysql driver library
```bash
go get github.com/go-redis/redis
```