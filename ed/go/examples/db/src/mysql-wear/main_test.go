package main

import (
	mw "github.com/clinc/mysql-wear"
	"testing"
)

var db *mw.DB

func init() {
	// suppose you have some helper code for initializing db connection
	db = dbtest.InitDB()
}

func TestUser2020CRUD(t *testing.T) {
	user := NewUser2020()
	// Fill in struct properties here, especially the ID/PK field

	if err := user.Insert(db); err != nil {
		log.Fatalf("insert failed: %v", err)
	}

	// Make sure we can get the newly inserted object
	user2, err := GetUser2020(db, user.ID)
	if err != nil {
		log.Fatalf("fail get item %s: %v", user.ID, err)
	}
	if user2 == nil {
		t.Fatalf("Didnt find newly inserted row with ID %s", user.ID)
	}
	// Make some changes to user here

	if err := user.Update(db); err != nil {
		log.Fatalf("id (%s): update failed: %v", user.ID, err)
	}

	// Make sure those changes took effect
	user3, err := GetUser2020(db, user.ID)
	if err != nil {
		log.Fatalf("fail get item %s: %v", user3.ID, err)
	}
	if user3 == nil {
		t.Fatalf("Missing row 3 ID %s", user3.ID)
	}

	// Compare props

}

func TestUserEmail2020CRUD(t *testing.T) {
	email := NewUserEmail2020()
	// Fill in struct properties here, especially the ID/PK field

	if err := email.Insert(db); err != nil {
		log.Fatalf("insert failed: %v", err)
	}

	// Make sure we can get the newly inserted object
	email2, err := GetUserEmail2020(db, email.ID)
	if err != nil {
		log.Fatalf("fail get item %s: %v", email.ID, err)
	}
	if email2 == nil {
		t.Fatalf("Didnt find newly inserted row with ID %s", email.ID)
	}
	// Make some changes to email here

	if err := email.Update(db); err != nil {
		log.Fatalf("id (%s): update failed: %v", email.ID, err)
	}

	// Make sure those changes took effect
	email3, err := GetUserEmail2020(db, email.ID)
	if err != nil {
		log.Fatalf("fail get item %s: %v", email3.ID, err)
	}
	if email3 == nil {
		t.Fatalf("Missing row 3 ID %s", email3.ID)
	}

	// Compare props

}
