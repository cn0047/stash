package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/clinc/mysql-wear/sqlq"

	mw "github.com/clinc/mysql-wear"
)

const (
	CONN_STR = "dbu:dbp@tcp(xmysql:3306)/test?charset=utf8"
)

type TestMysql struct {
	ID   int
	Code int
}

type UserProfile struct {
	ID        int `mw:"pk"`
	FirstName string
	LastName  string
	Tags      []string
	CreatedAt time.Time
}

func main() {
	f1()
	f4b()
	f4c()
	f5()
	f7()
	f8()
	f9()
	f10()
	f11()
	j1()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getDB() *mw.DB {
	dbc, err := sql.Open("mysql", CONN_STR)
	checkErr(err)
	db := mw.New(dbc)
	return db
}

func f1() {
	db := getDB()
	v, err := db.Count(&TestMysql{})
	checkErr(err)
	fmt.Printf("\nf1. got count: %d\n", v)
}

func f2() {
	fmt.Println("\nf2.", mw.GenerateModel(&UserProfile{}, "up")) // panic: Missing primary key for table (user_profile)
}

func f3() {
	fmt.Println("\nf3.", mw.GenerateModelTest(&UserProfile{}, "upt")) // panic: Missing primary key for table (user_profile)
}

func f4() {
	u1 := &UserProfile{ID: 1, FirstName: "James", LastName: "Bond"}
	db := getDB()
	db.MustInsert(u1) // panic: Error 1146: Table 'test.user_profile' doesn't exist
}

func f4b() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("\nf4b. ok")
		}
	}()

	t := &TestMysql{ID: 1, Code: 200}
	db := getDB()
	db.MustInsert(t)
}

func f4c() {
	t := &TestMysql{ID: 1, Code: 200}
	db := getDB()
	r, err := db.Insert(t)
	fmt.Printf("\nf4c. Got err: %#v\n", err)

	if r != nil {
		id, err := r.LastInsertId()
		fmt.Printf("\n\t Got id: %#v, error: %#v\n", id, err)
	}
}

func f5() {
	var data []TestMysql
	db := getDB()
	db.MustSelect(&data, sqlq.Columns("id", "code"),
		sqlq.Limit(10), sqlq.GreaterThan("code", 100),
	)
	fmt.Printf("\nf5. Got data:\n%#v\n", data)
}

func f6() {
	var data []TestMysql
	db := getDB()
	db.MustSelect(&data, sqlq.Columns("id", "code + 1000")) // panic: unrecognized column (code + 1000)
	fmt.Printf("\nf6. Got data:\n%#v\n", data)
}

func f6a() {
	var data []TestMysql
	db := getDB()
	db.MustSelect(&data, sqlq.Columns("id", "avg(code)")) // panic: unrecognized column (avg(code))
	fmt.Printf("\nf6a. Got data:\n%#v\n", data)
}

func f7() {
	t := &TestMysql{}
	db := getDB()
	found := db.MustGet(t, sqlq.Equal("code", 200))
	fmt.Printf("\nf7. Found: %v, Got data: %#v\n", found, t)
}

func f8() {
	t := &TestMysql{}
	db := getDB()
	found := db.MustGet(t, sqlq.Equal("code", 200))

	if !found {
		return
	}

	t.Code = 201
	if err := db.Update(t); err != nil {
		fmt.Printf("\nf8. Got update error 1: %#v\n", err)
	}

	t.Code = 200
	if err := db.Update(t); err != nil {
		fmt.Printf("\nf8. Got update error 2: %#v\n", err)
	}
}

func f9() {
	t := &TestMysql{ID: 3, Code: 501}
	db := getDB()
	db.MustInsert(t) // panic: Please pass a struct pointer to parseModel

	if err := db.Delete(t); err != nil {
		fmt.Printf("\nf9. Got delete error: %#v\n", err)
	}
}

func f10() {
	db := getDB()
	count := db.MustCount(&TestMysql{}, sqlq.LessThan("code", 1000))
	fmt.Printf("\nf10. Got count: %d\n", count)
}

func f11() {
	fmt.Println("\nf11.", mw.GenerateSchema(&UserProfile{}))
}

func j1() {
	question, answers, err := QuestionByIdAndAnswers("2")
	fmt.Printf("\nj1. err: %#v \n", err)
	fmt.Printf("\t answers: %v \n", answers)
	fmt.Printf("\t question: %v \n", question)
	fmt.Printf("\t question.Answers: %v \n", question.Answers[0])
	fmt.Printf("\t question.Answers: %v \n", question.Answers[1])
}

type Question struct {
	ID      int
	Body    string
	Answers []Answer `mw:"-"`
}

type Answer struct {
	ID         int
	QuestionID int
	Body       string
}

func QuestionByIdAndAnswers(id string) (*Question, []Answer, error) {
	db := getDB()

	question := &Question{ID: 2}
	found := db.MustGet(question)
	if !found {
		return question, nil, fmt.Errorf("question not found")
	}

	answers := make([]Answer, 0)
	db.MustSelect(&answers, sqlq.Equal("question_id", question.ID))
	question.Answers = answers

	return question, answers, nil
}

// -------------------------------------------- //
// AUTO GENERATED
// -------------------------------------------- //

func NewUserProfile() *UserProfile {
	return &UserProfile{}
}

func GetUserProfile(db *mw.DB, id int) (*UserProfile, error) {
	up := &UserProfile{ID: id}
	found, err := db.Get(up)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, nil
	}
	return up, nil
}

func (up *UserProfile) Insert(db *mw.DB) error {
	//up.Created = time.Now().UTC()
	//up.Updated = time.Now().UTC()
	res, err := db.Insert(up)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("fail get last id: %v", err)
	}
	up.ID = int(id)
	return nil
}

func (up *UserProfile) Update(db *mw.DB) error {
	//up.Updated = time.Now().UTC()
	if err := db.Update(up); err != nil {
		return err
	}
	return nil
}

func (up *UserProfile) Delete(db *mw.DB) error {
	if err := db.Delete(up); err != nil {
		return err
	}
	return nil
}

// -------------------------------------------- //
// END AUTO GENERATED
// -------------------------------------------- //
