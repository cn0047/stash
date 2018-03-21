Products
-

## Description

This is simple products web server.

## Usage

````
# Install dependencies:
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
    go get -u github.com/labstack/echo/...
'

# Run web server:
docker run -it --rm -p 8080:8080 -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
    go run src/github.com/ncube/products/main.go
'

# Check:
curl -XGET 'http://localhost:8080/products'
curl -XGET 'http://localhost:8080/products/mTd3lb'

# Test:
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
    cd src/github.com/ncube/products/controller && go test -v
'
````
