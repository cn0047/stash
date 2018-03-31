Products REST-API
-

[![Go Report Card](https://goreportcard.com/badge/github.com/cn007b/pra)](https://goreportcard.com/report/github.com/cn007b/pra)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/63e723d6c3f04982b7dc3270c84fc288)](https://www.codacy.com/app/cn007b/pra?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=cn007b/pra&amp;utm_campaign=Badge_Grade)
[![Maintainability](https://api.codeclimate.com/v1/badges/2cbdbfb4284431703c47/maintainability)](https://codeclimate.com/github/cn007b/pra/maintainability)

## Description

This is simple products web server.
<br>All products placed in `products.zip` archive.
<br>All Products will be extracted and served through REST API.

Endpoints:

* `GET /products` - return all known products.
* `GET /products/<product-id>` - return particular product.

## Usage

````
# Install dependencies:
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
    go get -u github.com/labstack/echo/...
'

# Run web server:
docker run -it --rm -p 8080:8080 -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
    go run src/github.com/app/products/main.go
'

# Check:
curl -XGET 'http://localhost:8080/products'
curl -XGET 'http://localhost:8080/products/mTd3lb'

# Test:
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
    cd src/github.com/app/products/ && go test -v ./...
'
````
