package dynamodb

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
	"strconv"
	"time"
)

type MovieItem struct {
	Year   int
	Title  string
	Plot   string
	Rating float64
}

type HotDataItem struct {
	Key    string `json:"key"`
	Val    string `json:"val"`
	ValInt int    `json:"val_int"`
}

func Run(cfg *aws.Config) {
	svc := dynamodb.New(session.New(), cfg)
	//movies(svc)
	//hotData(svc)
	//putHotDataItem(svc)
	//putHotDataItems(svc)
	getHotDataItem(svc, os.Args[1])
}

func movies(svc *dynamodb.DynamoDB) {
	createTableIfNotExists(svc)
	itemsCount(svc)
	putMovieItem(svc)
	getMovieItem(svc)
}

func hotData(svc *dynamodb.DynamoDB) {
	fmt.Printf("hotdata table exsists: %v\n", isTableExists(svc, "hotdata"))
	putHotDataItem(svc)
}

func putHotDataItems(svc *dynamodb.DynamoDB) {
	valInt, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Got error %w:", err)
	}

	for i := int64(0); i < 100; i++ {
		item := HotDataItem{
			Key:    os.Args[1] + strconv.FormatInt(time.Now().UnixNano()+i, 10),
			Val:    os.Args[2] + strconv.FormatInt(i, 10),
			ValInt: valInt,
		}
		putItem(svc, item, "hotdata")
	}
}

func putHotDataItem(svc *dynamodb.DynamoDB) {
	valInt, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Got error %w:", err)
	}

	item := HotDataItem{
		Key:    os.Args[1],
		Val:    os.Args[2],
		ValInt: valInt,
	}
	putItem(svc, item, "hotdata")
}

func isTableExists(svc *dynamodb.DynamoDB, tableNameToFind string) bool {
	tablesList, err := svc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		fmt.Println("Got error %w:", err.Error())
	}
	if tablesList == nil || tablesList.TableNames == nil {
		fmt.Println("Got nil output")
		return false
	}

	for _, tableName := range tablesList.TableNames {
		if *tableName == tableNameToFind {
			return true
		}
	}

	return false
}

func createTableIfNotExists(svc *dynamodb.DynamoDB) {
	if !isTableExists(svc, "Movies") {
		createTable(svc)
	}
}

func createTable(svc *dynamodb.DynamoDB) {
	input := &dynamodb.CreateTableInput{
		// see: https://docs.aws.amazon.com/sdkforruby/api/Aws/DynamoDB/Types/AttributeValue.html
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Year"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("Title"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("Year"), KeyType: aws.String("HASH")},   // partition key
			{AttributeName: aws.String("Title"), KeyType: aws.String("RANGE")}, // sort key
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String("Movies"),
	}

	_, err := svc.CreateTable(input)
	if err != nil {
		fmt.Println("Got error calling CreateTable:", err.Error())
		os.Exit(1)
	}
}

func itemsCount(svc *dynamodb.DynamoDB) {
	info, err := svc.DescribeTable(&dynamodb.DescribeTableInput{TableName: aws.String("Movies")})
	if err != nil {
		fmt.Println("Got error %w:", err.Error())
	}

	if info != nil && info.Table != nil {
		v := info.Table.ItemCount
		fmt.Printf("Found count in table: %v \n", *v)
	}
}

func putMovieItem(svc *dynamodb.DynamoDB) {
	item := MovieItem{
		Year:   2015,
		Title:  "The Big New Movie",
		Plot:   "Nothing happens at all.",
		Rating: 0.0,
	}
	putItem(svc, item, "Movies")
}

func putItem(svc *dynamodb.DynamoDB, item interface{}, tableName string) {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Got error marshalling new movie item:", err.Error())
		os.Exit(1)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	res, err := svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Result: %+v\n", res.String())
}

func getMovieItem(svc *dynamodb.DynamoDB) {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Movies"),
		Key: map[string]*dynamodb.AttributeValue{
			"Year": {
				N: aws.String("2015"),
			},
			"Title": {
				S: aws.String("The Big New Movie"),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Found movie: ", result)
}

func getHotDataItem(svc *dynamodb.DynamoDB, key string) {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("hotdata"),
		Key: map[string]*dynamodb.AttributeValue{
			"key": {
				S: aws.String(key),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Found item: ", result)
}
