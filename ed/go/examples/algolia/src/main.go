package main

import (
	"fmt"
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"time"
)

func main() {
	client := algoliasearch.NewClient("821NZIDT2Z", "01dcbdda46a6a787bd84a0ddfd069e35")
	index := client.InitIndex("my2")

	//f1(index)
	//f2(index)
	//f3(index)
	//f4(index)
	f5(index)
}

func f1(index algoliasearch.Index) {
	objects := []algoliasearch.Object{
		{"type": "tracking log", "time": time.Now()},
		{"objectID": "the main tracking log", "type": "tracking log", "time": time.Now()},
		{"objectID": "the update example", "type": "update example", "time": time.Now()},
	}

	res, err := index.AddObjects(objects)
	fmt.Printf("\n f1: %+v, %+v", res, err)
}

// f2 - AddObjects with already existing id -> update will be performed.
func f2(index algoliasearch.Index) {
	objects := []algoliasearch.Object{
		{"objectID": "the main tracking log", "type": "tracking-log", "time": time.Now()},
	}

	res, err := index.AddObjects(objects)
	fmt.Printf("\n f2: %+v, %+v", res, err)
}

// f3 - no update will be performed for object without objectID.
func f3(index algoliasearch.Index) {
	objects := []algoliasearch.Object{
		{"type": "update tracking log", "time": time.Now()}, // won't be added
		{"objectID": "the update example", "type": "update-example", "time": time.Now()},
	}

	res, err := index.UpdateObjects(objects)
	fmt.Printf("\n f3: %+v, %+v", res, err)
}

func f4(index algoliasearch.Index) {
	objects := []algoliasearch.Object{
		{"objectID": "the update example", "time": "recently"},
	}

	res, err := index.PartialUpdateObjects(objects)
	fmt.Printf("\n f4: %+v, %+v", res, err)
}

func f5(index algoliasearch.Index) {
	objects := []algoliasearch.Object{
		//{"type": "at1", "time": time.Now()}, // error: no objectID
		{"objectID": "at2", "type": "at2", "time": time.Now()},
	}

	res, err := index.PartialUpdateObjects(objects)
	fmt.Printf("\n f5: %+v, %+v", res, err)
}
