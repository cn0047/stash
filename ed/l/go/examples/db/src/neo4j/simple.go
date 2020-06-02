package main

import (
	"fmt"
	"time"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func main() {
	d := getDriver()
	defer d.Close()
	s := getSession(d)

	//f1(s)
	f2(s)
}

func getDriver() neo4j.Driver {
	dsn := ""
	dsn = "bolt://xneo4j:7687"
	//dsn = "bolt+routing://xneo4j:7687"
	user := "neo4j"
	pwd := "test"
	realm := ""

	useConsoleLogger := func(level neo4j.LogLevel) func(config *neo4j.Config) {
		return func(config *neo4j.Config) {
			config.Log = neo4j.ConsoleLogger(level)
		}
	}
	l := neo4j.ERROR // ERROR | DEBUG
	driver, err := neo4j.NewDriver(dsn, neo4j.BasicAuth(user, pwd, realm), useConsoleLogger(l))
	if err != nil {
		panic(fmt.Errorf("[getDriver] ERR: %v", err))
	}

	return driver
}

func getSession(driver neo4j.Driver) neo4j.Session {
	s, err := driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		panic(fmt.Errorf("[getSession] ERR: %v", err))
	}

	return s
}

func f2(session neo4j.Session) {
	params := map[string]interface{}{"code": "007"}
	result, err := session.Run(`MATCH (n:Person {code: $code}) RETURN n;`, params)
	if err != nil {
		panic(fmt.Errorf("ERR-3 : %v", err))
	}
	if result.Next() {
		r := result.Record().GetByIndex(0)
		fmt.Printf("🎾 %#v \n", r) // &neo4j.nodeValue{id:21, labels:[]string{"Person"}, props:map[string]interface {}{"active":true, "code":"007", "name":"James Bond"}}
	}
}

func f1(session neo4j.Session) {
	r, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		c := time.Now().Unix()
		result, err := transaction.Run(
			`CREATE (p:Person) SET p.code = $c, p.name = $n RETURN p.code + ', from node ' + id(p)`,
			map[string]interface{}{"c": c, "n": fmt.Sprintf("agent-%d", c)},
		)
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().GetByIndex(0), nil
		}

		return nil, result.Err()
	})
	if err != nil {
		fmt.Printf("[f1] ERR: %v", err)
	}
	fmt.Printf("🎾 %#v \n", r)
}

type Person struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Active bool   `json:"active"`
}
