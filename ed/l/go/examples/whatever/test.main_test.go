package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("🚧")
	r := m.Run()
	fmt.Println("🏁")
	os.Exit(r)

}

func TestX(t *testing.T) {
	t.Run("testCase1", func(t *testing.T) {
		t.Fatal()
	})

	t.Run("testCase2", func(t *testing.T) {
		t.FailNow()
	})
}

func TestY(t *testing.T) {
	t.Run("testCase1", func(t *testing.T) {
		t.Logf("🎾")
	})

	t.Run("testCase2", func(t *testing.T) {
		t.Logf("🔴")
	})
}
