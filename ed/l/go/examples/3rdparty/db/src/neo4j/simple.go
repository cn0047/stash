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

	exists, err := IsAccountExists(s, "main1")
	if err != nil {
		panic(err)
	}
	if !exists {
		err := createAccountTX(s, "main1")
		if err != nil {
			panic(err)
		}
	}
	for {
		incrementAccountBalance(s, "main1")
	}
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

func createAccountTX(session neo4j.Session, id interface{}) error {
	r, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		code := time.Now().Unix()
		result, err := tx.Run(
			`CREATE (a:Account) SET a.id = $id, a.balance = 0`,
			map[string]interface{}{"id": id},
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
		return fmt.Errorf("failed to create account, error: %w", err)
	}

	fmt.Printf("Account created, reslut: %+v \n", r)

	return nil
}

func IsAccountExists(session neo4j.Session, ID interface{}) (bool, error) {
	params := map[string]interface{}{"id": ID}
	result, err := session.Run(`MATCH (a:Account {id: $id}) RETURN a;`, params)
	if err != nil {
		return false, fmt.Errorf("failed to find account, error: %w", err)
	}

	if result.Next() {
		r := result.Record().GetByIndex(0)
		v, ok := r.(neo4j.Node)
		if !ok {
			return false, fmt.Errorf("got non neo4j.Node result")
		}

		if v.Props()["id"] != ID {
			return false, fmt.Errorf("got non expected ID")
		}

		return true, nil
	}

	return false, nil
}

func incrementAccountBalance(session neo4j.Session, ID interface{}) error {
	r, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		q := `
			MATCH (a:Account {id: $id})
			SET a.balance = a.balance + 1
			WITH a CALL apoc.util.sleep(1000)
			MATCH (ac:Account {balance: a.balance})
			RETURN ac
		`
		result, err := tx.Run(q, map[string]interface{}{"id": ID})
		if err != nil {
			return nil, fmt.Errorf("failed to incremented account balance, error: %w", err)
		}

		_, err = result.Summary()
		if err != nil {
			return nil, fmt.Errorf("failed to get summary, error: %w", err)
		}

		if result.Next() {
			r := result.Record().GetByIndex(0)
			v, ok := r.(neo4j.Node)
			if !ok {
				return false, fmt.Errorf("got non neo4j.Node result")
			}

			return v.Props()["balance"], nil
		}

		return nil, fmt.Errorf("failed to get result, error: %w", result.Err())
	})
	if err != nil {
		return fmt.Errorf("failed to perform TX. to incremented account balance, error: %w", err)
	}

	fmt.Printf("Account balance incremented, reslut: %+v \n", r)

	return nil
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
