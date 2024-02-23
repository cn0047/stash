package main

import (
	"log"
)

func main() {
	a, err := NewApp(WithVal("Key", "MyKey"))
	if err != nil {
		log.Printf("failed to create new app, err: %v\n", err)
		return
	}

	log.Printf("App.Key: %s\n", a.GetKey())
}

type App struct {
	Key string
}

func (a *App) GetKey() string {
	return a.Key
}

type AppOptionFunc func(s *App) error

func NewApp(options ...AppOptionFunc) (*App, error) {
	a := &App{}
	for _, option := range options {
		err := option(a)
		if err != nil {
			return nil, err
		}
	}

	return a, nil
}

func WithVal(key string, val any) AppOptionFunc {
	return func(a *App) error {
		// @TODO: Finish next line.
		log.Printf("set to key: %s, val: %v\n", key, val)
		return nil
	}
}
