Products
-

## Description

This is simple products web server.

JSON endpoints:

* `GET /products` - return all products known.

* `GET /products/<product-id>` - return particular product.

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


## Improvements

- Echo webserver could have been configured to log each request to the terminal (access log)
- Using GOPATH as the base dir as you have for unzipping can cause issues since gopath can
  configure more than one dir.
- Although creative, usually a better way to deal with this sort of json data
  would have been to first put it into a struct or map then send it to requesting
  clients. Your approach will load files from disk each request instead of from memory.
  Reading the products files into memory first then simply using echo's c.JSON()
  function to serialize to JSON would have simplified your product functions.

## Errors

- Given the GetAllProducts function writes status 200 at the top, the error condition
  will not actually write 500 status
- The test do not pass without first having run the webserver to load the data dir.
