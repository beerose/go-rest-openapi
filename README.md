**TL;DR command list**

    git clone https://github.com/gbrayhan/microservices-go
    cd microservices-go
    cp config.json.example config.json
    docker-compose up  --build  -d

### Build and run image of docker

```bash
docker-compose up  --build  -d
```

### Swagger Implementation

```bash
swag init -g infrastructure/rest/routes/routes.go
```

To visualize the swagger documentation on local use

http://localhost:8080/v1/swagger/index.html

### Unit test command

```bash
# run recursive test
go test  ./test/unit/...
# clean go test results in cache
go clean -testcache
```

### Lint inspection of go

```bash
golangci-lint run ./...
```
