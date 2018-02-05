GO
-

go1.9

````
$GOPATH

go run --work ed/go/examples/hw.go
````

````go
import (
    "fmt"
    "os"
)

var power int
power = 9000
fmt.Printf("It's over %d\n", power)

var power int = 9000
power := 9000

if (name == "Goku" && power > 9000) || (name == "gohan" && power < 4000) {
    print("super Saiyan")
    println("it's over 9000!")
}

if len(os.Args) != 2 {
    os.Exit(1)
}
fmt.Println("It's over", os.Args[1])

//
defer file.Close()
````

````go
for i := 0; i < 100; i++ {
}
````

#### Function Declaration

````go
func log(message string) { }
func add(a int, b int) int { }
func power(name string) (int, bool) { }

value, exists := power("goku")
_, exists := power("goku")
````

#### Structure

````go
type Saiyan struct {
    Name string
    Power int
}

goku := Saiyan {
  Name: "Goku",
  Power: 9000,
}

goku := Saiyan{}

goku := Saiyan{Name: "Goku"}
goku.Power = 9000

goku := Saiyan{"Goku", 9000}
````

#### Array

````go
var scores [10]int
scores[0] = 339
scores := [4]int{9001, 9333, 212, 33}

scores = append(scores, 5)

powers := make([]int, len(saiyans))
for index, saiyan := range saiyans {
    powers[index] = saiyan.Power
}
````

#### Map

````go
total := len(lookup)
delete(lookup, "goku")
````

#### Interface

````go
type Logger interface {
    Log(message string)
}

type Server struct {
    logger Logger
}
````

#### Goroutine

#### Channel
