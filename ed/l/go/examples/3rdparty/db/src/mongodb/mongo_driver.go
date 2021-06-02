package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()

	u := "mongodb://dbu:dbp@localhost:27017/test"
	collectionName := "test"
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(u))
	if err != nil {
		panic(fmt.Errorf("failed to connect to mongo, err: %w", err))
	}
	conn := c.Database("test")

	//findOneByCode(ctx, conn, collectionName, 202)
	//findByCodeWithTimeout(ctx, conn, collectionName, 202)
	findByCodeWithManualTimeout(ctx, conn, collectionName, 202)
}

func wait(d time.Duration, silent bool) {
	done := time.After(d)
	fmt.Printf("%v ", time.Now().Format(time.Kitchen))
	for {
		select {
		case <-done:
			if !silent {
				fmt.Printf(" %v\n", time.Now().Format(time.Kitchen))
			}
			return
		default:
			if !silent {
				fmt.Printf(".")
				time.Sleep(10 * time.Second)
			}
		}
	}
}

func findByCodeWithManualTimeout(ctx context.Context, db *mongo.Database, collection string, code int) {
	opts := options.Find().SetNoCursorTimeout(true)
	cur, err := db.Collection(collection).Find(ctx, bson.M{"code": bson.M{"$in": bson.A{code}}}, opts)
	if err != nil {
		panic(fmt.Errorf("failed to perform find, err: %w", err))
	}
	defer cur.Close(ctx)
	go func() {
		wait(2*time.Minute, true)
		fmt.Printf("manual cursor timeout\n")
		cur.Close(ctx)
	}()

	for cur.Next(ctx) {
		var result = bson.M{}
		err = cur.Decode(&result)
		if err != nil {
			panic(fmt.Errorf("failed to perform decode, err: %w", err))
		}
		fmt.Printf("doc: %+v \n", result)
		wait(1*time.Minute, false)
	}
	if err = cur.Err(); err != nil {
		panic(fmt.Errorf("got cursor err: %w", err))
	}
}

func findByCodeWithTimeout(ctx context.Context, db *mongo.Database, collection string, code int) {
	ctx2, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	opts := options.Find().SetNoCursorTimeout(true)
	cur, err := db.Collection(collection).Find(ctx2, bson.M{"code": bson.M{"$in": bson.A{code}}}, opts)
	if err != nil {
		panic(fmt.Errorf("failed to perform find, err: %w", err))
	}
	defer cur.Close(ctx2)

	for cur.Next(ctx2) {
		var result = bson.M{}
		err = cur.Decode(&result)
		if err != nil {
			panic(fmt.Errorf("failed to perform decode, err: %w", err))
		}
		fmt.Printf("doc: %+v \n", result)
		wait(1*time.Minute, false)
	}
	if err = cur.Err(); err != nil {
		panic(fmt.Errorf("got cursor err: %w", err))
	}
}

func findOneByCode(ctx context.Context, db *mongo.Database, collection string, code int) {
	opts := options.FindOne().SetProjection(bson.M{"code": 1})
	res := db.Collection(collection).FindOne(ctx, bson.M{"code": code}, opts)
	if res.Err() != nil {
		panic(fmt.Errorf("failed to perform find, err: %w", res.Err()))
	}

	bsonMap := bson.M{}
	err := res.Decode(bsonMap)
	if err != nil {
		panic(fmt.Errorf("failed to perform decode, err: %w", err))
	}

	fmt.Printf("🎾 %+v \n", bsonMap)
}
