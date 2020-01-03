package main

import (
	"database/sql"
	"fmt"
	"time"

	mw "github.com/clinc/mysql-wear"
	"github.com/clinc/mysql-wear/sqlq"
)

const (
	DB_CONN_STR = "dbu:dbp@tcp(xmysql:3306)/test?charset=utf8"
)

func main() {
	step4()
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func DB() *mw.DB {
	dbc, err := sql.Open("mysql", DB_CONN_STR)
	check(err)
	db := mw.New(dbc)
	return db
}

type User2020 struct {
	ID string `mw:"pk" json:"id"`
	// Don't ever pass back the password via json response
	PasswordHash string `json:"-"`
	// This is used to be able to destroy all user sessions in a login-less env
	// aka using jwt not a session server
	TokenVersion int `json:"-"`
	// Could use for membership payments. Most iOS app users already have this
	StripeCustomerID string `json:"-"`

	Status string `json:"status"` // active|locked|blocked or banned?|etc

	Created time.Time `mw:"-" json:"created"`
	Updated time.Time `mw:"-" json:"updated"`

	Emails []UserEmail2020 `mw:"join"`
}

type UserEmail2020 struct {
	// Primary Key
	Sha1    string `mw:"pk"`
	UserID  string `json:"user_id"`
	Email   string `json:"email"`
	Primary bool   `json:"primary"` // Primary aka default email

	Created time.Time `mw:"-" json:"created"`
	Updated time.Time `mw:"-" json:"updated"`

	VerifiedOn time.Time `mw:"-" json:"verified_on"` // If/when the user has verified this email. Only needed non-oauth
}

func step1() {
	u := &User2020{}
	fmt.Println(mw.GenerateModel(u, "user"))
	fmt.Println(mw.GenerateModelTest(u, "user"))
	fmt.Println(mw.GenerateSchema(u))

	ue := &UserEmail2020{}
	fmt.Println(mw.GenerateModel(ue, "email"))
	fmt.Println(mw.GenerateModelTest(ue, "email"))
	fmt.Println(mw.GenerateSchema(ue))
}

func step2() {
	db := DB()

	u := &User2020{ID: "u2"}
	found := db.MustGet(u)
	if !found {
		u.PasswordHash = "u2pwd"
		u.PasswordHash = "u2pwd"
		u.TokenVersion = 2
		u.StripeCustomerID = "tripe2"
		u.Status = "ok2"
		u.Created = time.Now()
		u.Updated = time.Now()
		db.MustInsert(u)
	}

	ue := &UserEmail2020{Sha1: "7F550A9F4C44173A37664D938F1355F0F92A47A7"}
	found = db.MustGet(ue)
	if !found {
		ue.UserID = u.ID
		ue.Email = "user2@email.com"
		ue.Primary = false
		ue.Created = time.Now()
		ue.Updated = time.Now()
		ue.VerifiedOn = time.Now()
		db.MustInsert(ue)
	}
}

func step3() {
	db := DB()

	u := &User2020{ID: "u2"}
	found := db.MustGet(u)
	fmt.Println(found)
	if found {
		ue := &UserEmail2020{Sha1: "7F550A9F4C44173A37664D938F1355F0F92A47A7"}
		db.MustGet(ue)
		fmt.Printf("\n Got user: %v", u)
		fmt.Printf("\n Got user email: %v", ue)
	}
}

func step4() {
	db := DB()

	u := make([]User2020, 0)
	err := db.Select(
		&u,
		sqlq.Join(&UserEmail2020{}, "user2020.id = user_email2020.user_id", "email"),
		sqlq.Equal("id", "u2"),
	)
	check(err)

	fmt.Printf("\n Got user's emails: %v", u[0].Emails)
}

// -------------------------------------------- //
// AUTO GENERATED
// -------------------------------------------- //

func NewUser2020() *User2020 {
	return &User2020{}
}

func GetUser2020(db *mw.DB, id string) (*User2020, error) {
	user := &User2020{ID: id}
	found, err := db.Get(user)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, nil
	}
	return user, nil
}

func (user *User2020) Insert(db *mw.DB) error {
	user.Created = time.Now().UTC()
	user.Updated = time.Now().UTC()
	_, err := db.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (user *User2020) Update(db *mw.DB) error {
	user.Updated = time.Now().UTC()
	if err := db.Update(user); err != nil {
		return err
	}
	return nil
}

func (user *User2020) Delete(db *mw.DB) error {
	if err := db.Delete(user); err != nil {
		return err
	}
	return nil
}

func NewUserEmail2020() *UserEmail2020 {
	return &UserEmail2020{}
}

func GetUserEmail2020(db *mw.DB, id string) (*UserEmail2020, error) {
	email := &UserEmail2020{Sha1: id}
	found, err := db.Get(email)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, nil
	}
	return email, nil
}

func (email *UserEmail2020) Insert(db *mw.DB) error {
	email.Created = time.Now().UTC()
	email.Updated = time.Now().UTC()
	_, err := db.Insert(email)
	if err != nil {
		return err
	}
	return nil
}

func (email *UserEmail2020) Update(db *mw.DB) error {
	email.Updated = time.Now().UTC()
	if err := db.Update(email); err != nil {
		return err
	}
	return nil
}

func (email *UserEmail2020) Delete(db *mw.DB) error {
	if err := db.Delete(email); err != nil {
		return err
	}
	return nil
}

// -------------------------------------------- //
// END AUTO GENERATED
// -------------------------------------------- //
