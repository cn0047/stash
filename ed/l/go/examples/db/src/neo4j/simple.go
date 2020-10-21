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
	defer s.Close()

	//createPersonTX(s)
	//createPerson(s)
	//matchPersonByCode(s, "007")
	//bookmarkCreatePerson(d, s)
	incrementAccountBalance(s)
}

func getDriver() neo4j.Driver {
	dsn := ""
	dsn = "bolt://xneo4j:7687"
	//dsn = "bolt+routing://xneo4j:7687"
	user := "neo4j"
	pwd := "1" // test|1
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

func getSession(driver neo4j.Driver, bookmarks ...string) neo4j.Session {
	s, err := driver.Session(neo4j.AccessModeWrite, bookmarks...)
	if err != nil {
		panic(fmt.Errorf("[getSession] ERR: %v", err))
	}

	return s
}
func incrementAccountBalance(session neo4j.Session) {
}

func bookmarkCreatePerson(d neo4j.Driver, s neo4j.Session) {
	bookmark, code := createPersonTX(s)
	s2 := getSession(d, bookmark)
	defer s2.Close()
	matchPersonByCodeTX(s2, code)
}

func matchPersonByCodeTX(session neo4j.Session, code interface{}) {
	r, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		params := map[string]interface{}{"code": code}
		result, err := tx.Run(`MATCH (n:Person {code: $code}) RETURN n;`, params)
		if err != nil {
			panic(fmt.Errorf("ERR-5 : %v", err))
		}
		if !result.Next() {
			return nil, result.Err()
		}

		r := result.Record().GetByIndex(0)
		return r, nil
	})
	if err != nil {
		panic(fmt.Errorf("ERR-4 : %v", err))
	}

	fmt.Printf("🆃🎾 %#v \n", r)
}

func matchPersonByCode(session neo4j.Session, code interface{}) {
	params := map[string]interface{}{"code": code}
	result, err := session.Run(`MATCH (n:Person {code: $code}) RETURN n;`, params)
	if err != nil {
		panic(fmt.Errorf("ERR-3 : %v", err))
	}
	if result.Next() {
		r := result.Record().GetByIndex(0)
		fmt.Printf("🎾 %#v \n", r) // &neo4j.nodeValue{id:21, labels:[]string{"Person"}, props:map[string]interface {}{"active":true, "code":"007", "name":"James Bond"}}
	}
}

func createPerson(session neo4j.Session) {
	c := time.Now().Unix()
	result, err := session.Run(
		`CREATE (p:Person) SET p.code = $c, p.name = $n RETURN p.code + ', from node ' + id(p)`,
		map[string]interface{}{"c": c, "n": fmt.Sprintf("agent-%d", c)},
	)
	if err != nil {
		fmt.Printf("[f1b] ERR-1: %v", err)
		return
	}

	r, err := result.Consume()
	if err != nil {
		fmt.Printf("[f1b] ERR-2: %v", err)
		return
	}

	fmt.Printf("🎾 %#v \n", r)
}

func createPersonTX(session neo4j.Session) (bookmark string, code interface{}) {
	r, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		code := time.Now().Unix()
		result, err := tx.Run(
			`CREATE (p:Person) SET p.code = $c, p.name = $n RETURN 'code:' + p.code + ', id:' + id(p)`,
			map[string]interface{}{"c": code, "n": fmt.Sprintf("agent-%d", code)},
		)
		if err != nil {
			return nil, err
		}

		summary, err := result.Summary()
		if err != nil {
			return nil, err
		}
		fmt.Printf("summary: %+v \n", summary)

		if !result.Next() {
			return nil, result.Err()
		}
		fmt.Printf("result: %#v \n", result.Record().GetByIndex(0))

		return code, nil
	})
	if err != nil {
		fmt.Printf("[f1] ERR: %v", err)
	}

	code = r
	fmt.Printf("🆃🎾 code: %#v \n", code)

	bookmark = session.LastBookmark()
	fmt.Printf("🆃🎾 bookmark: %#v \n", bookmark)

	return bookmark, code
}

type Person struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Active bool   `json:"active"`
}
