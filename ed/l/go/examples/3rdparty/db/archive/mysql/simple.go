/*

create table test_mysql (
	id int key, code int
);
insert into test_mysql values
	(1, 200);

create table question (
	id int auto_increment key,
	body varchar(100)
);
insert into question values
	(1, "question 1"),
	(2, "question 2");

create table answer (
	id int auto_increment key,
	question_id int,
	body varchar(100)
);
insert into answer values
	(1, 1, "answer 1 to q 1"),
	(2, 2, "answer 1 to q 2"),
	(3, 2, "answer 2 to q 2");

*/

package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const (
	CONN_STR =                                   //
	"dbu:dbp@tcp(xmysql:3306)/test?charset=utf8" // docker-compose
	// "root:root@tcp(xmysql:3306)/?charset=utf8" // docker
	// "dbu:dbp@tcp(xmysql:3306)/test?charset=utf8" // docker
	// "dbu:dbp@tcp(172.17.0.4:3306)/test?charset=utf8" // k8sg
)

func main() {
	hw()
	// f0()
	// f0a()
	// f1()
	// f2()
	// f2b()
	// web()
	// j1()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func f4() {
	db, err := sql.Open("mysql", CONN_STR)

	rows, err := db.Query("CALL getCountry(?)", "ua")
	checkErr(err)

	for rows.Next() {
		id := -1
		name := ""
		err = rows.Scan(&id, &name)
		checkErr(err)
		fmt.Printf("Got country with code: %v and name: %v \n", id, name)
	}
}

func hw() {
	db, err := sql.Open("mysql", CONN_STR)
	checkErr(err)

	rows, err := db.Query("SELECT ?", 204)
	checkErr(err)

	for rows.Next() {
		v := -1
		err = rows.Scan(&v)
		checkErr(err)
		fmt.Printf("Got value: %d \n", v)
	}
}

func f2() {
	db, err := sql.Open("mysql", CONN_STR)
	checkErr(err)

	rows, err := db.Query("SELECT * FROM test_mysql")
	checkErr(err)

	for rows.Next() {
		id := -1
		code := -1
		err = rows.Scan(&id, &code)
		checkErr(err)
		fmt.Printf("Got row with id: %d, and code: %d \n", id, code)
	}
}

func f2b() {
	db, err := sql.Open("mysql", CONN_STR)
	checkErr(err)

	rows, err := db.Query("SELECT * FROM tmp")
	checkErr(err)

	for rows.Next() {
		n := -1
		m := ""
		err = rows.Scan(&n, &m)
		checkErr(err)
		fmt.Printf("Got row with n: %d, and m: %v \n", n, m)
	}
}

func f1() {
	db, err := sql.Open("mysql", CONN_STR)
	checkErr(err)

	stmt, err := db.Prepare("INSERT IGNORE INTO test_mysql VALUES (?, ?)")
	checkErr(err)

	res, err := stmt.Exec(2, 204)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("Inserted id:", id)
}

func f0a() {
	db, err := sql.Open("mysql", CONN_STR)
	checkErr(err)

	_, err = db.Exec("CREATE DATABASE test3")
	checkErr(err)
}

func f0() {
	db, err := sql.Open("mysql", CONN_STR)
	checkErr(err)

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS test_mysql (id int key, code int)")
	checkErr(err)

	_, er2 := stmt.Exec()
	checkErr(er2)

	stmt2, err := db.Prepare("REPLACE test_mysql VALUES (?, ?)")
	checkErr(err)

	_, err3 := stmt2.Exec(1, 200)
	checkErr(err3)
}

func web() {
	db, err := sql.Open("mysql", CONN_STR)
	checkErr(err)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM test_mysql")
		checkErr(err)

		w.Write([]byte("<table>"))
		w.Write([]byte("<tr><td>id</td><td>code</td></tr>"))
		for rows.Next() {
			id := -1
			code := -1
			err = rows.Scan(&id, &code)
			checkErr(err)
			w.Write([]byte(fmt.Sprintf("<tr> <td>%d</td> <td>%d</td> </tr>", id, code)))
		}
		w.Write([]byte("</table>"))
	})

	p := ":8080"
	fmt.Printf("Open: http://localhost%s\n", p)
	http.ListenAndServe(p, nil)
}

func j1() {
	question, answers, err := QuestionByIdAndAnswers("2")
	fmt.Printf("\nj1. err: %#v \n", err)
	fmt.Printf("\t answers: %v \n", answers)
	fmt.Printf("\t question: %v \n", question)
	fmt.Printf("\t question.Answers: %v \n", *question.Answers[0])
	fmt.Printf("\t question.Answers: %v \n", *question.Answers[1])
}

type Question struct {
	ID      int
	Body    string
	Answers []*Answer
}

type Answer struct {
	ID         int
	QuestionID int
	Body       string
}

func QuestionByIdAndAnswers(id string) (*Question, []*Answer, error) {
	db, err := sql.Open("mysql", CONN_STR)
	checkErr(err)

	query := `
		SELECT q.id, q.body, a.id, a.question_id, a.body
		FROM question AS q
		JOIN answer AS a ON q.id = a.question_id
		WHERE q.id = ?
	`
	rows, err := db.Query(query, id)
	checkErr(err)

	question := &Question{}
	for rows.Next() {
		answer := &Answer{}
		err = rows.Scan(
			&question.ID,
			&question.Body,
			&answer.ID,
			&answer.QuestionID,
			&answer.Body,
		)
		checkErr(err)
		question.Answers = append(question.Answers, answer)
	}

	return question, question.Answers, nil
}
