Sets Calculator
-

<br>⚠️ DISCLAIMER: Due to lack of time this calculator works only with 2 sets per 1 command.
<br>⚠️ DISCLAIMER: By default command DIF produces symmetric difference for 2 sets.

## Test

````sh
docker run -it --rm -v $PWD:/app -w /app cn007b/go sh -c '
  go test -v ./...
'
````

## Build

````sh
docker run -it --rm -v $PWD:/app -w /app cn007b/go sh -c '
  go build -ldflags="-w -s" -o scalc main.go
'
````

## Run

````sh
docker run -it --rm -v $PWD:/app -w /app cn007b/go ./scalc [ INT a.txt b.txt ]
docker run -it --rm -v $PWD:/app -w /app cn007b/go ./scalc [ DIF b.txt c.txt ]
docker run -it --rm -v $PWD:/app -w /app cn007b/go ./scalc [ SUM a.txt c.txt ]
docker run -it --rm -v $PWD:/app -w /app cn007b/go ./scalc [ SUM a.txt [ SUM b.txt c.txt ] ]
docker run -it --rm -v $PWD:/app -w /app cn007b/go ./scalc [ SUM [ DIF a.txt b.txt ] [ INT b.txt c.txt ] ]
docker run -it --rm -v $PWD:/app -w /app cn007b/go ./scalc [ SUM [ DIF a.txt [ SUM a.txt c.txt ] ] [ INT b.txt c.txt ] ]
````
